package service

import (
	"github.com/widodragon/goshop/database"
	"github.com/widodragon/goshop/migration"
)

type UserService interface {
	SaveUser(user migration.User) error
	SaveCredit(user migration.CreditCard) error
	UpdateUser(user migration.User) error
	DeleteUser(user migration.User) error
	FindAllUser() []migration.User
}

type userService struct {
	userDatabase database.UserDatabase
}

func NewUser(userDatabase database.UserDatabase) UserService {
	return &userService{
		userDatabase: userDatabase,
	}
}

func (db *userService) SaveUser(user migration.User) error {
	err := db.userDatabase.SaveUser(user)
	return err
}

func (db *userService) SaveCredit(user migration.CreditCard) error {
	err := db.userDatabase.SaveCredit(user)
	return err
}

func (db *userService) UpdateUser(user migration.User) error {
	db.userDatabase.UpdateUser(user)
	return nil
}

func (db *userService) DeleteUser(user migration.User) error {
	db.userDatabase.DeleteUser(user)
	return nil
}

func (db *userService) FindAllUser() []migration.User {
	return db.userDatabase.FindAllUser()
}
