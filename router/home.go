package router

import (
	"thefireseed/model"
	"thefireseed/utils"

	"github.com/gin-gonic/gin"
	_model "github.com/lazyfury/go-web-template/model"
	"gorm.io/gorm"
)

func homePage(c *gin.Context) {
	page, size := _model.GetPagingParams(c)
	apps := model.DB.GetObjectsOrEmpty(&[]model.AppModel{}, map[string]interface{}{})

	apps.Paging(page, size, func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at desc")
	})

	utils.Render(c, "homePage", utils.UserParam{
		"apps":   apps.Result.List.(*[]model.AppModel),
		"paging": apps.Pagination,
	}, utils.DefaultSEO().SetTitle("home"), utils.Banner{
		Title: `"火种计划"是一个很牛x的`,
		Tips:  "为了人类的为了，进步，科技xxxxx",
	})
}
