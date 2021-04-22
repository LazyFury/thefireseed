package router

import (
	"shareInviteCode/model"
	"shareInviteCode/utils"

	"github.com/gin-gonic/gin"
	_model "github.com/lazyfury/go-web-template/model"
	"github.com/lazyfury/go-web-template/response"
	"gorm.io/gorm"
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
	}, utils.DefaultSEO().SetTitle(app.AppName))
}

func activityDetail(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		response.Error("获取活动内容错误")
	}

	activity := &model.ActivityModel{}
	if err := model.DB.GetObjectOrNotFound(activity, map[string]interface{}{
		"code": code,
	}); err != nil {
		response.Error(err)
	}

	app := &model.AppModel{}
	if err := model.DB.GetObjectOrNotFound(app, map[string]interface{}{
		"code": activity.AppCode,
	}); err != nil {
		response.Error(err)
	}

	page, size := _model.GetPagingParams(c)
	codes := model.DB.GetObjectsOrEmpty(&[]model.CodeModel{}, map[string]interface{}{
		"activity_code": activity.Code,
	}, func(db *gorm.DB) *gorm.DB {
		return db.Order("used asc,created_at desc")
	})
	codes.Paging(page, size)

	list := codes.Result.List.(*[]model.CodeModel)
	for i, invite := range *list {
		str := invite.InviteCode
		strLen := len(str)
		if strLen >= 8 {
			b := []byte{}
			for i := 0; i < (strLen - 6); i++ {
				b = append(b, []byte("*")...)
			}
			str = str[:3] + string(b) + str[strLen-3:]
		}
		(*list)[i].InviteCode = str
	}

	utils.Render(c, "activityDetail", utils.UserParam{
		"activity": activity,
		"app":      app,
		"codes":    list,
		"paging":   codes.Pagination,
	})
}
