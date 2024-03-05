package auth

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
