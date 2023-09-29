package http

import (
	"errors"

	"go-chat/db"
	"go-chat/internal/storage"
	"go-chat/pkg/models"
	"go-chat/pkg/utils"

	"github.com/charmbracelet/log"
	v "github.com/cohesivestack/valgo"
	"github.com/gofiber/fiber/v2"
	"github.com/oklog/ulid/v2"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Name            string `json:"name"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

func Login(c *fiber.Ctx) error {
	body := user{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	val := v.
		Is(v.String(body.Username, "username").Not().Blank()).
		Is(v.String(body.Password, "password").MinLength(8))

	if !val.Valid() {
		return c.Status(fiber.ErrBadRequest.Code).JSON(val.Error())
	}

	user, err := db.GetUserByUsername(body.Username, c.Context())
	if err != nil {
		return c.Status(fiber.ErrUnauthorized.Code).JSON(fiber.Map{
			"err": "username or password is incorrect",
		})
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return c.Status(fiber.ErrUnauthorized.Code).JSON(fiber.Map{
			"err": "username or password is incorrect",
		})
	}

	err = createSession(c, user.Username)
	if err != nil {
		return utils.InternalErr(c, errors.New("user created but login failed"))
	}

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	// i use 1 minute session expration time so for now its okey
	// TODO implement the logout
	return fiber.ErrNotImplemented
}

func Register(c *fiber.Ctx) error {
	body := user{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	val := v.
		Is(v.String(body.Name, "name").Not().Blank()).
		Is(v.String(body.Username, "username").Not().Blank()).
		Is(v.String(body.Password, "password").MinLength(8)).
		Is(v.String(body.PasswordConfirm, "password_confirm").EqualTo(body.Password, "Passwords must be same"))

	if !val.Valid() {
		return c.Status(fiber.ErrBadRequest.Code).JSON(val.Error())
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("bcrypt hash error", "err", err)
		return utils.InternalErr(c, err)
	}

	u := models.User{
		ID:       ulid.Make().String(),
		Name:     body.Name,
		Username: body.Username,
		Password: string(hash),
	}

	userExist, err := db.CheckUsername(u.Username, c.Context())
	if err != nil {
		return utils.InternalErr(c, err)
	}
	if userExist {
		return c.Status(fiber.ErrConflict.Code).JSON(fiber.Map{
			"err": "username already exists",
		})
	}

	_, err = db.CreateUser(&u, c.Context())
	if err != nil {
		log.Error("registerHandler", "err", err)
		return utils.InternalErr(c, err)
	}

	// TODO find better way to login after registeration and do better error handling
	createSession(c, u.Username)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": u.ID,
	})
}

func createSession(c *fiber.Ctx, u string) error {
	log := log.WithPrefix("SESSION")
	userSess, err := storage.Session.Get(c)
	if err != nil {
		log.Error("user session get error", "err", err)
	}

	userSess.Set("user", u)
	err = userSess.Save()
	if err != nil {
		log.Error("user session save error", "err", err)
	}
	return err
}
