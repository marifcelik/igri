package auth

import (
	"errors"
	"fmt"
	"time"

	"go-chat/models"
	"go-chat/storage"
	"go-chat/utils"

	"github.com/charmbracelet/log"
	v "github.com/cohesivestack/valgo"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler interface {
	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	createSession(c *fiber.Ctx, u string) error
}

type authHandler struct {
	repo *authRepo
}

func NewAuthHandler(repo *authRepo) AuthHandler {
	return &authHandler{repo: repo}
}

func (h *authHandler) Login(c *fiber.Ctx) error {
	body := UserLoginDTO{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	val := v.
		Is(v.String(body.Username, "username").Not().Blank()).
		Is(v.String(body.Password, "password").Not().Blank().MinLength(8))

	if !val.Valid() {
		return c.Status(fiber.ErrBadRequest.Code).JSON(val.Error())
	}

	user, err := h.repo.GetUserByUsername(body.Username, c.Context())
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(fiber.ErrUnauthorized.Code).JSON(fiber.Map{
				"err": "username or password is incorrect",
			})
		}
		log.Error(fmt.Errorf("get user by username error: %w", err))
		return utils.InternalErr(c, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return c.Status(fiber.ErrUnauthorized.Code).JSON(fiber.Map{
			"err": "username or password is incorrect",
		})
	}

	err = h.createSession(c, user.Username)
	if err != nil {
		return utils.InternalErr(c, errors.New("username and password are correct but login failed"))
	}

	return c.JSON(user)
}

func (h *authHandler) Logout(c *fiber.Ctx) error {
	// i use 1 minute session expration time so for now its okey
	// TODO implement the logout
	return fiber.ErrNotImplemented
}

func (h *authHandler) Register(c *fiber.Ctx) error {
	body := UserRegisterDTO{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	val := v.
		Is(v.String(body.Name, "name").Not().Blank()).
		Is(v.String(body.Username, "username").Not().Blank()).
		Is(v.String(body.Password, "password").MinLength(8).Not().Blank()).
		Is(v.String(body.PasswordConfirm, "passwordConfirm").EqualTo(body.Password, "Passwords must be same"))

	if !val.Valid() {
		return c.Status(fiber.ErrBadRequest.Code).JSON(val.Error())
	}

	userExist, err := h.repo.CheckUsername(body.Username, c.Context())
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		log.Error(fmt.Errorf("check username error: %w", err))
		return utils.InternalErr(c, err)
	}
	if userExist {
		return c.Status(fiber.ErrConflict.Code).JSON(fiber.Map{
			"err": "username already exists",
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("bcrypt hash error", "err", err)
		return utils.InternalErr(c, err)
	}

	u := &models.User{
		Name:     body.Name,
		Username: body.Username,
		Password: string(hash),
		M: models.M{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	id, err := h.repo.CreateUser(u, c.Context())
	if err != nil {
		log.Error("user create error", "err", err)
		return utils.InternalErr(c, err)
	}
	log.Print("user created", "user", u)

	// TODO find better way to login after registeration and do better error handling
	// FIX sometimes the returned authorization header is: Bearer {\n "err": "username or password is incorrect"\n}
	err = h.createSession(c, u.Username)
	if err != nil {
		return utils.InternalErr(c, errors.New("user created but login failed"))
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":       id,
		"username": u.Username,
	})
}

func (h *authHandler) createSession(c *fiber.Ctx, u string) error {
	log := log.WithPrefix("SESSION")
	userSess, err := storage.Session.Get(c)
	if err != nil {
		log.Error("user session get error", "err", err)
		return err
	}

	userSess.Set("user", u)
	err = userSess.Save()
	if err != nil {
		log.Error("user session save error", "err", err)
	}
	return err
}
