package Models

type User struct {
	ID       uint   `gorm:"column:id;primary_key;not null"`
	Email    string `gorm:"column:email;not null;unique" json:"email" binding:"required"`
	Password string `gorm:"column:password;not null" json:"password" binding:"required"`
	Salt     string `gorm:"column:salt;not null" json:"salt"`
}
