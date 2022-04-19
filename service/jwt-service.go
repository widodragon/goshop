package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken(user string, admin string) (string, error)
	ValidateToken(jwtToken string) (*jwt.Token, error)
}

type MyCustomClaims struct {
	User  string `json:"user"`
	Admin string `json:"admin"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: "JWT_SECRET",
		issuer:    "test",
	}
}

func (service *jwtService) GenerateToken(user string, admin string) (string, error) {
	claims := MyCustomClaims{
		user,
		admin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    service.issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(service.secretKey))
	return ss, err
}

func (service *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Signing method validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret signing key
		return []byte([]byte(service.secretKey)), nil
	})

}
