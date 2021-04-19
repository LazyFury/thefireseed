package router

import (
	"net/http"
	"shareInviteCode/model"

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
	c.HTML(http.StatusOK, "home/index.html", map[string]interface{}{
		"apps": apps.Result.List.(*[]model.AppModel),
	})
}
