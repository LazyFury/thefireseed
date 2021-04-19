package router

import (
	"net/http"
	"shareInviteCode/model"

	"github.com/gin-gonic/gin"
	"github.com/lazyfury/go-web-template/response"
)

func appDetail(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		response.Error("获取App内容错误")
	}
	app := &model.AppModel{}
	if err := model.DB.GetObjectOrNotFound(app, map[string]interface{}{
		"code": code,
	}); err != nil {
		response.Error(err)
	}
	c.HTML(http.StatusOK, "home/app/detail.html", map[string]interface{}{
		"app": app,
	})
}
