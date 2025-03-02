package Models

type PasswordUpdate struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Pin      string `json:"pin" binding:"required"`
}
