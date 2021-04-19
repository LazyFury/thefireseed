package router

import (
	"shareInviteCode/model"
	"shareInviteCode/utils"

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

	utils.Render(c, "appDetail", map[string]interface{}{
		"app": app,
	})
}
