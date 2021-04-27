package controller

import (
	"io"
	"net/http"
	"thefireseed/model"

	"github.com/gin-gonic/gin"
	"github.com/lazyfury/go-web-template/controller"
	"github.com/lazyfury/go-web-template/response"
)

type CodeLogsController struct {
	controller.Controller
}

func NewCodeLogsController() *CodeLogsController {
	return &CodeLogsController{
		controller.Controller{
			DB:    model.DB,
			Model: &model.CodeCopyLogModel{},
			Auth:  Auth,
		},
	}
}

func (d *CodeLogsController) Install(g *gin.RouterGroup, path string) {
	controller.Install(g, d, path)
}

func (d *CodeLogsController) Add(c *gin.Context) {
	user := GetUserOrLogin(c)

	logs := &model.CodeCopyLogModel{}

	if err := c.ShouldBind(logs); err != nil {
		if err == io.EOF {
			panic("没有传入参数，请使用post json传入参数")
		}
		panic(err)
	}
	logs.UserCode = user.Code
	if err := logs.Validator(); err != nil {
		panic(err)
	}

	logs.SetCode()

	if err := model.DB.Create(logs).Error; err != nil {
		panic(err)
	}
	c.JSON(http.StatusCreated, response.JSON(response.StatusCreated, "", logs))
}
