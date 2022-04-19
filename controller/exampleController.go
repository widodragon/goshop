package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type userTest struct {
	Name   string `form:"name"`
	Alamat string `form:"alamat"`
}

type ExampleController interface {
	SaveUser(ctx *gin.Context) gin.H
}

type exampleController struct {
	users []userTest
}

func NewExample() ExampleController {
	return &exampleController{}
}

func (c *exampleController) SaveUser(ctx *gin.Context) gin.H {
	var exampleForm userTest
	err := ctx.ShouldBind(&exampleForm)
	if err != nil {
		return gin.H{
			"error": err.Error(),
			"users": c.users,
		}
	}
	fmt.Println(exampleForm.Name)
	c.users = append(c.users, exampleForm)
	return gin.H{
		"error": nil,
		"users": c.users,
	}
}
