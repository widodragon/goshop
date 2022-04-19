package entity

type Token struct {
	User     string `form:"user"`
	Admin    string `form:"admin"`
	Password string `form:"password"`
}
