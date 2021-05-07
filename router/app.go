package router

import (
	"fmt"

	"github.com/lazyfure/thefireseed/controller"
	"github.com/lazyfure/thefireseed/model"
	"github.com/lazyfure/thefireseed/utils"

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
	codeModel := &model.CodeModel{}
	copyModel := &model.CodeCopyLogModel{}
	usedModel := &model.CodeUsedLogModel{}
	user := controller.GetUserOrEmpty(c)
	codes := model.DB.GetObjectsOrEmpty(&[]model.ShowCodeModal{}, map[string]interface{}{
		"activity_code": activity.Code,
	}, func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at desc")
	}, codeModel.Joins,
		// 当前用户是否复制过
		func(db *gorm.DB) *gorm.DB {
			return db.Joins(
				fmt.Sprintf(
					"left join (select 1 `copied`,`invite_id` `copy_invite_id` from `%s` where `user_code`='%s') clip_t on clip_t.`copy_invite_id`=`%s`.`code`",
					copyModel.TableName(),
					user.Code,
					codeModel.TableName(),
				),
			)
		},
		func(db *gorm.DB) *gorm.DB {
			return db.Joins(
				fmt.Sprintf(
					"left join (select 1 `used`,`invite_id` `used_invite_id` from `%s` where `user_code`='%s') used_t on used_t.`used_invite_id`=`%s`.`code`",
					usedModel.TableName(),
					user.Code,
					codeModel.TableName(),
				),
			)
		},
	)
	codes.Paging(page, size)

	utils.Render(c, "activityDetail", utils.UserParam{
		"activity": activity,
		"app":      app,
		"codes":    codeModel.Result(codes.Result.List).(*[]model.ShowCodeModal),
		"paging":   codes.Pagination,
	}, utils.DefaultSEO().SetTitle(fmt.Sprintf("%s %s", activity.ActivityName, app.AppName)))
}
