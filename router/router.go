package router

import (
	"shareInviteCode/controller"
	"shareInviteCode/middleware"
	"shareInviteCode/utils"

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

	// 登录
	app.GET("/login", login)
	app.GET("/reg", reg)
	app.GET("/logout", controller.NewUserController().Logout)

}

func login(c *gin.Context) {
	utils.Render(c, "login")
}
func reg(c *gin.Context) {
	utils.Render(c, "reg")
}
