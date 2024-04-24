package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"go-chat/models"
	st "go-chat/storage"
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
	createSession(r *http.Request, u string)
}

type authHandler struct {
	repo *AuthRepo
}

func NewAuthHandler(repo *AuthRepo) AuthHandler {
	return &authHandler{repo: repo}
}

func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {
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
		utils.ErrResp(w, http.StatusBadRequest, val.Error())
		return
	}

	user, err := h.repo.GetUserByUsername(body.Username, r.Context())
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			utils.JsonResp(w, utils.M{"err": "username or password is incorrect"}, http.StatusUnauthorized)
			return
		}
		log.Error(fmt.Errorf("get user by username error: %w", err))
		utils.ErrResp(w, http.StatusInternalServerError, err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		utils.JsonResp(w, utils.M{"err": "username or password is incorrect"}, http.StatusUnauthorized)
		return
	}

	h.createSession(r, user.Username)

	utils.JsonResp(w, user)
}

func (h *authHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// i use 1 minute session expration time so for now its okey
	// TODO implement the logout
	userSession := st.Session.Exists(r.Context(), "user")
	if !userSession {
		utils.ErrResp(w, http.StatusNotAcceptable)
		return
	}

	err := st.Session.Destroy(r.Context())
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
		utils.JsonResp(w, val.Error(), http.StatusBadRequest)
		return
	}

	userExist, err := h.repo.CheckUsername(body.Username, r.Context())
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		log.Error(fmt.Errorf("check username error: %w", err))
		utils.InternalErrResp(w, err)
		return
	}
	if userExist {
		utils.JsonResp(w, utils.M{"err": "username already exists"}, http.StatusConflict)
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

	// TODO find better way to login after registeration and do better error handling
	// FIX sometimes the returned authorization header is: Bearer {\n "err": "username or password is incorrect"\n}
	h.createSession(r, user.Username)

	utils.JsonResp(w, utils.M{"id": id, "username": user.Username}, http.StatusCreated)
}

func (h *authHandler) createSession(r *http.Request, username string) {
	log := clog.WithPrefix("SESSION")
	log.Info("before put", "stat", st.Session.Status(r.Context()))
	st.Session.Put(r.Context(), "user", username)
	fmt.Printf("create storage.Session.Keys(r.Context()): %v\n", st.Session.Keys(r.Context()))
	log.Info("after put", "stat", st.Session.Status(r.Context()))
}
