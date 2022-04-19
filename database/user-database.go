package database

import (
	"github.com/widodragon/goshop/migration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserDatabase interface {
	SaveUser(user migration.User) error
	SaveCredit(credit migration.CreditCard) error
	UpdateUser(user migration.User)
	DeleteUser(user migration.User)
	FindAllUser() []migration.User
}

type database struct {
	connection *gorm.DB
}

func NewDatabase() UserDatabase {
	dsn := "host=127.0.0.1 user=admin password=24021994 dbname=golang port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&migration.User{}, &migration.CreditCard{})
	return &database{
		connection: db,
	}
}

func (data *database) SaveUser(user migration.User) error {
	result := data.connection.Create(&user)
	return result.Error
}

func (data *database) SaveCredit(credit migration.CreditCard) error {
	result := data.connection.Create(&credit)
	return result.Error
}

func (data *database) UpdateUser(user migration.User) {
	data.connection.Updates(&user)
}

func (data *database) DeleteUser(user migration.User) {
	data.connection.Delete(&user)
}

func (data *database) FindAllUser() []migration.User {
	var users []migration.User
	data.connection.Preload("CreditCard").Find(&users)
	return users
}
