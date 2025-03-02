package DataBase

import (
	"MPT-Schedule/Models"
	"time"
)

func ClearPincodes() {
	//очищение просроченных паролей каждые 5 секунд
	for {
		DB.Where("ended < ?", time.Now()).Delete(&Models.PinCode{})
		time.Sleep(5 * time.Second)
	}
}
