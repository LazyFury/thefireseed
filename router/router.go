package router

import (
	"shareInviteCode/controller"
	"shareInviteCode/middleware"

	"github.com/gin-gonic/gin"
)

func Init(c *gin.RouterGroup) {
	c.Use(middleware.AuthOrNot)

	// api
	api := c.Group("/api")
	controller.NewAppController().Install(api, "/apps")
	controller.NewActivityController().Install(api, "/activities")
	controller.NewCodeController().Install(api, "/activities/:id/codes")
	controller.NewUserController().Install(api, "/users")

	// 前端模板
	app := c.Group("")
	app.GET("", homePage)
	app.GET("/app/:code", appDetail)
	app.GET("/activities/:code", activityDetail)

	// 登录
	app.GET("/login", login)
	app.GET("/reg", reg)
	app.GET("/logout", controller.NewUserController().Logout)

}
