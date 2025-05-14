package Models

import "time"

type PinCode struct {
	ID      uint      `gorm:"primary_key;AUTO_INCREMENT"`
	Email   string    `json:"email" binding:"required"`
	Pin     string    `json:"pin"`
	Created time.Time `json:"created,omitempty"` // timestamp without time zone
	Ended   time.Time `json:"ended,omitempty"`   // timestamp without time zone
}
