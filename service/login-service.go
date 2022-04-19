package service

import (
	"github.com/widodragon/goshop/entity"
	"golang.org/x/crypto/bcrypt"
)

type LoginService interface {
	Login(user string, admin string, currentPassword string) bool
	SaveUser(client entity.User) entity.User
	GetAllUser() []entity.User
	HashPassword(password string) (string, error)
}

type loginService struct {
	JWTService
	users []entity.User
}

func NewLoginService(JWTService JWTService) LoginService {
	return &loginService{
		JWTService: JWTService,
	}
}

func (login *loginService) Login(user string, admin string, currentPassword string) bool {
	users := login.GetAllUser()
	var password string
	if users != nil || len(users) != 0 {
		for _, data := range users {
			if data.User == user && data.Admin == admin {
				password = data.Password
			}
		}
	}
	isMatch := ValidatePassword(password, currentPassword)
	if isMatch == true {
		return true
	}
	return false
}

func ValidatePassword(passwordHash string, currentPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(currentPassword))
	return err == nil
}

func (login *loginService) GetAllUser() []entity.User {
	return login.users
}

func (login *loginService) SaveUser(client entity.User) entity.User {
	login.users = append(login.users, client)
	return client
}

func (login *loginService) HashPassword(password string) (string, error) {
	var message error
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 11)
	if err != nil {
		message = err
	}
	return string(hashPassword), message
}
