package controller

import (
	"shareInviteCode/model"

	"github.com/gin-gonic/gin"
	"github.com/lazyfury/go-web-template/controller"
	"github.com/lazyfury/go-web-template/response"
)

type CodeController struct {
	controller.Controller
}

func NewCodeController() *CodeController {
	return &CodeController{
		controller.Controller{
			DB:    model.DB,
			Model: &model.CodeModel{},
			Auth:  Auth,
		},
	}
}

func (d *CodeController) Install(g *gin.RouterGroup, path string) {
	controller.Install(g, d, path)
	router := g.Group(path)
	router.GET("/:id/used", d.SetUsed)
}

func (d *CodeController) SetUsed(c *gin.Context) {
	code := c.Param("id")
	if code == "" {
		response.Error("code参数错误")
	}

	if err := model.DB.Model(&model.CodeModel{}).Where(map[string]interface{}{
		"code": code,
	}).Update("used", 1).Error; err != nil {
		response.Error(err)
	}

	response.Error(response.JSONSuccess("操作成功", nil))
}
