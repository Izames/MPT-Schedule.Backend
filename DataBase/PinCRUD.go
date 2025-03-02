package DataBase

import (
	"MPT-Schedule/Models"
	"fmt"
	"time"
)

func CreatePin(email string, pincode string) string {
	var pin Models.PinCode
	pin.Pin = pincode
	pin.Email = email
	if result := DB.Create(&pin); result.Error != nil {
		fmt.Println(result.Error)
		return "ошибка сохранения пин кода в базу данных"
	}
	return ""
}
func FindPinByMail(mail string) (Models.PinCode, error) {
	var pin Models.PinCode
	err := DB.Where("email= ? AND ended > ?", mail, time.Now()).First(&pin).Error
	if err != nil {
		return pin, err
	}
	return pin, nil
}
