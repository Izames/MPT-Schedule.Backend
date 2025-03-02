package DataBase

import (
	"MPT-Schedule/Models"
	"log"
)

func SetDefault() {
	var user Models.User
	//проверка есть ли дефолт юзер в системе
	err := DB.Where("email=?", "geday2@mail.ru").First(&user).Error
	if err == nil {
		log.Println("дефолтный пользователь уже есть в системе")
	} else {
		CreateUser("geday2@mail.ru", "123456")
		log.Println("дефолтный пользователь создан")

	}
}
