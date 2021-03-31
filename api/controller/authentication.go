package controller

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func GenerateTokenWithName(name string) (string, error) {

	// 鍵となる文字列(ここではセキュリティを求めないので全部この鍵を用いる)
	secretKey := "secret"

	// トークンはjwtを使用
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": name,
	})
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	fmt.Printf("tokenString: %v", tokenString)

	return tokenString, nil

}
