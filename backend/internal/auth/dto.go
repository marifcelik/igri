package auth

import (
	"time"
)

type UserLoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegisterDTO struct {
	Name            string `json:"name"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type UserReturnDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
}
