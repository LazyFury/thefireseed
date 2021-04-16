package router

import (
	"shareInviteCode/controller"
	"shareInviteCode/middleware"

	"github.com/gin-gonic/gin"
)

func Init(c *gin.RouterGroup) {
	api := c.Group("/api", middleware.AuthOrNot)
	controller.NewAppController().Install(api, "/apps")
	controller.NewActivityController().Install(api, "/activities")
	controller.NewCodeController().Install(api, "/activities/:id/codes")
}
