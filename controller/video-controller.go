package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/widodragon/goshop/entity"
	"github.com/widodragon/goshop/service"
	"github.com/widodragon/goshop/validation"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) gin.H
	ShowPageHtml(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-wido", validation.TitleValidation)
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) gin.H {
	var video entity.Video
	err := ctx.BindJSON(&video)
	if err != nil {
		return gin.H{
			"error": err.Error(),
		}
	}
	err = validate.Struct(video)
	if err != nil {
		return gin.H{
			"error": err.Error(),
		}
	}
	c.service.Save(video)
	return gin.H{
		"error": err,
		"data":  video,
	}
}

func (c *controller) ShowPageHtml(ctx *gin.Context) {
	data := c.service.FindAll()
	passingData := gin.H{
		"title": "Halaman Utama",
		"data":  data,
	}
	ctx.HTML(http.StatusOK, "index.html", passingData)
}
