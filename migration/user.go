package migration

import "time"

type User struct {
	ID         uint64       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string       `json:"name" gorm:"type:varchar(200)"`
	Username   string       `json:"username" gorm:"type:varchar(10);UNIQUE"`
	Password   string       `json:"password" gorm:"type:varchar(200)"`
	CreditCard []CreditCard `gorm:"foreignKey:UserID"`
	CreateAt   time.Time    `json:"create_at" gorm:"autoCreateTime:nano"`
	UpdateAt   time.Time    `json:"update_at" gorm:"autoUpdateTime:nano"`
}

type CreditCard struct {
	Number string `json:"number" gorm:"type:varchar(200)"`
	UserID uint   `json:"id" gorm:"type:int(11)"`
}
