package JWT

import (
	"MPT-Schedule/Models"
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"time"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateJWT(user Models.User) (string, error) {
	tokenTTl, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTl)).Unix(),
	})
	return token.SignedString(privateKey)
}
