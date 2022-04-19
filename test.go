package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var mySigningKey []byte = []byte("JWT_SECRET")

func testing() {
	token, err := GenerateToken("wido", "wido24")
	if err != nil {
		fmt.Println(err.Error())
	}
	res, err := ValidateToken(token)
	fmt.Println(res.Raw)
	fmt.Println(token)
}

type MyCustomClaims struct {
	User  string `json:"user"`
	Admin string `json:"admin"`
	jwt.StandardClaims
}

func GenerateToken(user string, admin string) (string, error) {
	claims := MyCustomClaims{
		user,
		admin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	return ss, err
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Signing method validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret signing key
		return []byte(mySigningKey), nil
	})

}
