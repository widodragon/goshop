package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/widodragon/goshop/migration"
	"github.com/widodragon/goshop/service"
	"github.com/widodragon/goshop/validation"
)

type UserController interface {
	SaveUser(ctx *gin.Context) gin.H
	SaveCredit(ctx *gin.Context) gin.H
	// UpdateUser(ctx *gin.Context) error
	// DeleteUser(ctx *gin.Context) error
	FindAllUser() []migration.User
}

type userController struct {
	service service.UserService
}

func NewUser(service service.UserService) UserController {
	validate = validator.New()
	validate.RegisterValidation("is-wido", validation.TitleValidation)
	return &userController{
		service: service,
	}
}

func (c *userController) FindAllUser() []migration.User {
	return c.service.FindAllUser()
}

func (c *userController) SaveUser(ctx *gin.Context) gin.H {
	var user migration.User
	err := ctx.BindJSON(&user)
	if err != nil {
		return gin.H{
			"error": err.Error(),
		}
	}
	err = validate.Struct(user)
	if err != nil {
		return gin.H{
			"error": err.Error(),
		}
	}
	err = c.service.SaveUser(user)
	if err != nil {
		return gin.H{
			"error": err.Error(),
		}
	}
	return gin.H{
		"error": err,
		"data":  user,
	}
}

func (c *userController) SaveCredit(ctx *gin.Context) gin.H {
	var credit migration.CreditCard
	err := ctx.BindJSON(&credit)
	if err != nil {
		return gin.H{
			"error": err.Error(),
		}
	}
	err = validate.Struct(credit)
	if err != nil {
		return gin.H{
			"error": err.Error(),
		}
	}
	err = c.service.SaveCredit(credit)
	if err != nil {
		return gin.H{
			"error": err.Error(),
		}
	}
	return gin.H{
		"error": err,
		"data":  credit,
	}
}

// func (c *userController) ShowPageHtml(ctx *gin.Context) {
// 	data := c.service.FindAllUser()
// 	passingData := gin.H{
// 		"title": "Halaman Utama",
// 		"data":  data,
// 	}
// 	ctx.HTML(http.StatusOK, "index.html", passingData)
// }
