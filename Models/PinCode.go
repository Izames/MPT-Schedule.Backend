package Models

type PinCode struct {
	ID    uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Email string `json:"email" binding:"required"`
	Pin   string `json:"pin"`
}
