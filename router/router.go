package router

import (
	"thefireseed/controller"
	"thefireseed/middleware"

	"github.com/gin-gonic/gin"
)

func Init(c *gin.RouterGroup) {
	c.Use(middleware.AuthOrNot)

	// api
	api := c.Group("/api")
	controller.NewAppController().Install(api, "/apps")
	controller.NewActivityController().Install(api, "/activities")
	controller.NewCodeController().Install(api, "/codes")
	controller.NewCodeLogsController().Install(api, "/code-copy-logs")

	userController := controller.NewUserController()
	userController.Install(api, "/users")

	// 前端模板
	app := c.Group("")
	app.GET("", homePage)
	app.GET("/app/:code", appDetail)
	app.GET("/activities/:code", activityDetail)

	// 登录
	app.GET("/login", login)
	app.GET("/reg", reg)
	app.GET("/logout", userController.Logout)

}
