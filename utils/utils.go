package utils

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"nc_user.com/v1/config"
	model "nc_user.com/v1/model"
)

func Md5(input string) string {
	data := []byte(input)
	output := fmt.Sprintf("%s", md5.Sum(data))
	return output
}

func JwtGenerate(userID int, phone, email string) (string, error) {
	mySigningKey := []byte(config.Config.JWTSecret.JWTKey)

	// Create the Claims
	claims := model.UserClaims{
		UserID: userID,
		Phone:  phone,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(48 * time.Hour).Unix(), //unix timestamp
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(mySigningKey)
	return ss, nil
}
