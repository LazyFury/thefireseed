package router

import (
	"net/http"
	"shareInviteCode/controller"
	"shareInviteCode/middleware"

	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, templatePath string, data map[string]interface{}) {
	c.HTML(http.StatusOK, templatePath, map[string]interface{}{
		"data":   data,
		"header": "header",
		"footer": "footer",
	})
}

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
	render(c, "home/login.html", map[string]interface{}{})
}
func reg(c *gin.Context) {
	render(c, "home/reg.html", map[string]interface{}{})
}
