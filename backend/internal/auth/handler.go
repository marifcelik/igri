package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"go-chat/config"
	"go-chat/models"
	"go-chat/storage"
	"go-chat/utils"

	clog "github.com/charmbracelet/log"
	v "github.com/cohesivestack/valgo"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var log = clog.WithPrefix("AUTH")

type AuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	// isLoggedIn(r *http.Request, u string) bool
	// checkLoggedIn(r *http.Request, u string) bool
	createSession(r *http.Request, u string)
}

type authHandler struct {
	repo *AuthRepo
}

func NewAuthHandler(repo *AuthRepo) AuthHandler {
	return &authHandler{repo: repo}
}

func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	// TODO remove this sleep after testing
	if config.C.AppEnv == config.DevEnv {
		time.Sleep(time.Second * 2)
	}

	body := UserLoginDTO{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Warn("body parsing", "err", err)
		utils.ErrResp(w, http.StatusBadRequest)
		return
	}

	val := v.
		Is(v.String(body.Username, "username").Not().Blank()).
		Is(v.String(body.Password, "password").Not().Blank().MinLength(8))

	if !val.Valid() {
		utils.JsonResp(w, utils.M{
			"status":  "error",
			"message": "Validation error",
			"data":    val.Error(),
		}, http.StatusBadRequest)
		return
	}

	user, err := h.repo.GetUserByUsername(body.Username, r.Context())
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			utils.JsonResp(w, utils.M{"status": "error", "data": "username or password is incorrect"}, http.StatusUnauthorized)
			return
		}
		log.Error(fmt.Errorf("get user by username error: %w", err))
		utils.ErrResp(w, http.StatusInternalServerError, err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		utils.JsonResp(w, utils.M{"status": "error", "data": "username or password is incorrect"}, http.StatusUnauthorized)
		return
	}

	result := UserReturnDTO{}
	// TODO use CopyFields function after fix
	result.ID = user.ID.Hex()
	result.Name = user.Name
	result.Username = user.Username
	result.CreatedAt = user.CreatedAt

	h.createSession(r, user.ID.Hex())

	utils.JsonResp(w, utils.M{"status": "success", "data": result})
}

func (h *authHandler) Logout(w http.ResponseWriter, r *http.Request) {
	if !storage.Session.Exists(r.Context(), "user") {
		utils.ErrResp(w, http.StatusNotAcceptable)
		return
	}

	err := storage.Session.Destroy(r.Context())
	if err != nil {
		utils.ErrResp(w, http.StatusNotAcceptable, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	body := UserRegisterDTO{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Warn("body parsing", "err", err)
		utils.ErrResp(w, http.StatusBadRequest)
		return
	}

	val := v.
		Is(v.String(body.Name, "name").Not().Blank()).
		Is(v.String(body.Username, "username").Not().Blank()).
		Is(v.String(body.Password, "password").MinLength(8).Not().Blank()).
		Is(v.String(body.PasswordConfirm, "passwordConfirm").EqualTo(body.Password, "Passwords must be same"))

	if !val.Valid() {
		utils.JsonResp(w, utils.M{
			"status":  "error",
			"message": "Validation error",
			"data":    val.Error(),
		}, http.StatusBadRequest)
		return
	}

	userExist, err := h.repo.CheckUsername(body.Username, r.Context())
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		log.Error(fmt.Errorf("check username error: %w", err))
		utils.InternalErrResp(w, err)
		return
	}
	if userExist {
		utils.JsonResp(w, utils.M{"status": "error", "data": "username already exists"}, http.StatusConflict)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("bcrypt hash error", "err", err)
		utils.InternalErrResp(w, err)
		return
	}

	user := &models.User{
		Name:     body.Name,
		Username: body.Username,
		Password: string(hash),
		M: models.M{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	id, err := h.repo.CreateUser(user, r.Context())
	if err != nil {
		log.Error("user create error", "err", err)
		utils.InternalErrResp(w, err)
		return
	}
	log.Print("user created", "user", user)

	result := UserReturnDTO{}
	// TODO use CopyFields function after fix
	result.ID = id.Hex()
	result.Name = user.Name
	result.Username = user.Username
	result.CreatedAt = user.CreatedAt

	// if err = utils.CopyFields(user, &result); err != nil {
	// 	log.Error("copy fields error", "err", err)
	// 	utils.InternalErrResp(w, err)
	// 	return
	// }
	// fmt.Printf("after CopyFields function, result: %v\n", result)

	// TODO find better way to login after registeration and do better error handling
	// FIX sometimes the returned authorization header is: Bearer {\n "err": "username or password is incorrect"\n}
	h.createSession(r, user.ID.Hex())

	w.Header().Set("Location", fmt.Sprintf("/users/%s", id.Hex()))
	utils.JsonResp(w, utils.M{"status": "success", "data": result}, http.StatusCreated)
}

// TODO implement the isLoggedIn and checkLoggedIn
// func (h *authHandler) isLoggedIn(r *http.Request, u string) bool {
// 	user := storage.Session.GetString(r.Context(), "user")
// 	return user != ""
// }

// func (h *authHandler) checkLoggedIn(r *http.Request, u string) bool {
// 	user := storage.Session.GetString(r.Context(), "user")
// 	return user != "" && user == u
// }

func (h *authHandler) createSession(r *http.Request, userID string) {
	log := clog.WithPrefix("SESSION")
	log.Info("before put", "stat", storage.Session.Status(r.Context()))
	storage.Session.Put(r.Context(), "user", userID)
	fmt.Printf("create storage.Session.Keys(r.Context()): %v\n", storage.Session.Keys(r.Context()))
	log.Info("after put", "stat", storage.Session.Status(r.Context()))
}
