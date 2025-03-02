package DataBase

import (
	"MPT-Schedule/Models"
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func CreateUser(email string, password string) string {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		log.Fatal("ошибка генерации соли: ", err)
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("ошибка хэширования: ", err)
	}
	user := Models.User{
		Email:    email,
		Password: string(hashPassword),
		Salt:     base64.StdEncoding.EncodeToString(salt),
	}
	err = DB.Create(&user).Error
	if err != nil {
		log.Fatal("ошибка создания пользователя: ", err)
	}
	log.Printf("в систему был добавлен пользователь [ %s ]", user.Email)
	return "Successfully created user"
}
func FindUserByEmail(email string) Models.User {
	var user Models.User
	err := DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Println(err)
	}
	return user
}
func Update_user(user Models.User) string {
	salt := make([]byte, 16)  // Выберите подходящий размер соли
	_, err := rand.Read(salt) // Заполняем соль случайными байтами
	if err != nil {
		log.Fatal("Failed to generate salt: ", err)
	}
	// Хешируем пароль с солью
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash password: ", err)
	}
	// Сохраняем соль и хеш пароля в базу данных
	user.Password = string(hashedPassword)              // Хеш
	user.Salt = base64.StdEncoding.EncodeToString(salt) // Соль

	if result := DB.Model(&user).Updates(Models.User{Password: user.Password}).Error; result != nil {
		return "ошибка обновления пользователя"
	}
	return ""
}
