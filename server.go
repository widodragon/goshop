package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/widodragon/goshop/controller"
	"github.com/widodragon/goshop/database"
	"github.com/widodragon/goshop/middleware"
	"github.com/widodragon/goshop/service"
)

var (
	videoService      service.VideoService         = service.New()
	videoController   controller.VideoController   = controller.New(videoService)
	jwtService        service.JWTService           = service.NewJWTService()
	loginService      service.LoginService         = service.NewLoginService(jwtService)
	loginController   controller.LoginController   = controller.NewLoginController(loginService, jwtService)
	databaseUser      database.UserDatabase        = database.NewDatabase()
	userService       service.UserService          = service.NewUser(databaseUser)
	userController    controller.UserController    = controller.NewUser(userService)
	exampleController controller.ExampleController = controller.NewExample()
)

func CreateLoggerOutput() {
	f, _ := os.Create("logger/logger.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	CreateLoggerOutput()
	server := gin.New()
	server.Use(middleware.Logger())
	server.Static("/css", "./templates/css")
	server.StaticFile("/favicon.png", "./templates/assets/favicon.png")
	server.LoadHTMLGlob("templates/*.html")
	server.POST("/user", func(ctx *gin.Context) {
		res := loginController.SaveUser(ctx)
		if res["error"] != nil {
			ctx.JSON(http.StatusBadRequest, res)
		} else {
			ctx.JSON(http.StatusOK, res)
		}
	})
	server.GET("/login", func(ctx *gin.Context) {
		res := loginController.Login(ctx)
		if res["error"] != nil {
			ctx.JSON(http.StatusBadRequest, res)
		} else {
			ctx.JSON(http.StatusOK, res)
		}
	})
	server.POST("/example", func(ctx *gin.Context) {
		res := exampleController.SaveUser(ctx)
		if res["error"] != nil {
			ctx.JSON(http.StatusBadRequest, res)
		} else {
			ctx.JSON(http.StatusOK, res)
		}
	})
	server.POST("/user-database", func(ctx *gin.Context) {
		res := userController.SaveUser(ctx)
		if res["error"] != nil {
			ctx.JSON(http.StatusBadRequest, res["error"])
		} else {
			ctx.JSON(http.StatusOK, res)
		}
	})
	server.POST("/credit", func(ctx *gin.Context) {
		res := userController.SaveCredit(ctx)
		if res["error"] != nil {
			ctx.JSON(http.StatusBadRequest, res["error"])
		} else {
			ctx.JSON(http.StatusOK, res)
		}
	})
	server.GET("/user-database", func(ctx *gin.Context) {
		data := userController.FindAllUser()
		ctx.JSON(http.StatusOK, data)
	})
	apiGroup := server.Group("/api", middleware.JWTMiddleware())
	{
		apiGroup.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})
		apiGroup.POST("/video", func(ctx *gin.Context) {
			res := videoController.Save(ctx)
			if res["error"] != nil {
				ctx.JSON(http.StatusBadRequest, res)
			} else {
				ctx.JSON(http.StatusOK, res)
			}
		})
	}
	viewGroup := server.Group("/view")
	{
		viewGroup.GET("/", videoController.ShowPageHtml)
	}
	server.Run(":7777")
}
