package entity

type User struct {
	Name     string `json:"name" binding:"required"`
	User     string `json:"user" binding:"required"`
	Admin    string `json:"admin" binding:"required"`
	Password string `json:"password" binding:"required"`
}
