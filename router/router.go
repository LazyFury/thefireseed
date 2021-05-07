package router

import (
	"github.com/lazyfure/thefireseed/controller"
	"github.com/lazyfure/thefireseed/middleware"
	"github.com/lazyfure/thefireseed/utils"

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
	controller.NewCodeUsedController().Install(api, "/code-used-logs")

	userController := controller.NewUserController()
	userController.Install(api, "/users")

	// 前端模板
	app := c.Group("")
	app.GET("", homePage)
	app.GET("/app/:code", appDetail)
	app.GET("/activities/:code", activityDetail)

	app.GET("/fireseed", fireseed)
	app.GET("/about", about)
	// 登录
	app.GET("/login", login)
	app.GET("/reg", reg)
	app.GET("/logout", userController.Logout)

	app.GET("/profile", profile)
}

func about(c *gin.Context) {
	utils.Render(c, "home/about.html", utils.UserParam{})
}

func fireseed(c *gin.Context) {
	utils.Render(c, "home/seed.html")
}
