package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/widodragon/goshop/entity"
	"github.com/widodragon/goshop/service"
)

type LoginController interface {
	Login(ctx *gin.Context) gin.H
	SaveUser(ctx *gin.Context) gin.H
}

type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

func NewLoginController(loginService service.LoginService, jwtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (service *loginController) Login(ctx *gin.Context) gin.H {
	var user entity.Token
	var token string = ""
	var err error
	ctx.ShouldBindQuery(&user)
	isLogin := service.loginService.Login(user.User, user.Admin, user.Password)
	if isLogin {
		token, err = service.jwtService.GenerateToken(user.User, user.Admin)
		if err != nil {
			return gin.H{
				"token": nil,
				"err":   err,
			}
		}
		return gin.H{
			"token": token,
			"err":   err,
		}
	} else {
		return gin.H{
			"token": token,
			"err":   "password atau user tidak sesuai",
		}
	}
}
func (service *loginController) SaveUser(ctx *gin.Context) gin.H {
	var user entity.User
	err := ctx.BindJSON(&user)
	if err != nil {
		return gin.H{
			"error": err.Error(),
			"user":  user,
		}
	}
	user.Password, err = service.loginService.HashPassword(user.Password)
	service.loginService.SaveUser(user)
	return gin.H{
		"error": nil,
		"user":  user,
	}
}
