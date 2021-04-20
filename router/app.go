package router

import (
	"shareInviteCode/model"
	"shareInviteCode/utils"

	"github.com/gin-gonic/gin"
	_model "github.com/lazyfury/go-web-template/model"
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

	page, size := _model.GetPagingParams(c)
	activities := model.DB.GetObjectsOrEmpty(&[]model.ActivityModel{}, map[string]interface{}{
		"app_code": app.Code,
	})
	activities.Paging(page, size)

	utils.Render(c, "appDetail", utils.UserParam{
		"app":        app,
		"activities": activities.Result.List.(*[]model.ActivityModel),
		"paging":     activities.Pagination,
	})
}

func activityDetail(c *gin.Context) {

	utils.Render(c, "activityDetail", map[string]interface{}{})
}
