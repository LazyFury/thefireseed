package router

import (
	"shareInviteCode/controller"
	"shareInviteCode/middleware"
	"shareInviteCode/utils/layout"

	"github.com/gin-gonic/gin"
)

func Init(c *gin.RouterGroup) {
	// api
	api := c.Group("/api", middleware.AuthOrNot)
	controller.NewAppController().Install(api, "/apps")
	controller.NewActivityController().Install(api, "/activities")
	controller.NewCodeController().Install(api, "/activities/:id/codes")
	controller.NewUserController().Install(api, "/users")

	// 前端模板
	app := c
	app.GET("", homePage)
	app.GET("/app/:code", appDetail)

	// 登录
	app.GET("/login", login)
	app.GET("/reg", reg)

}

func login(c *gin.Context) {
	layout.New("login").Render(c)
}
func reg(c *gin.Context) {
	layout.New("reg").Render(c)
}
