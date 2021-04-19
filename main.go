package main

import (
	"shareInviteCode/config"
	"shareInviteCode/middleware"
	"shareInviteCode/model"
	"shareInviteCode/router"
	"shareInviteCode/utils"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	gwt "github.com/lazyfury/go-web-template"
	"github.com/lazyfury/go-web-template/response"
	"github.com/lazyfury/go-web-template/tools/template/layout"
)

func main() {
	// init
	app := gwt.New()

	// 连接数据库
	if err := model.DB.ConnectMysql(config.Global.Mysql.ToString()); err != nil {
		panic(err)
	}

	model.DB.AutoMigrate(
		&model.AppModel{}, &model.ActivityModel{}, &model.CodeModel{},
		&model.User{},
	)

	// 注册模版
	layout.InitBootstrap("templates", "*.html", utils.TemplateFuncs)
	app.SetHTMLTemplate(layout.Bootstrap)

	// 注册静态目录
	app.Use(static.Serve("/static", static.LocalFile("static", false)))

	// 注册路由

	router.Init(&app.RouterGroup)

	// 没有注册的路由，不走全局中间件，TODO:
	app.NoMethod(middleware.AuthOrNot, func(c *gin.Context) {
		response.Error(response.NoMethod)
	})
	app.NoRoute(middleware.AuthOrNot, func(c *gin.Context) {
		response.Error(response.NoRoute)
	})
	// 错误码配置
	response.RecoverErrHtml = true
	response.RecoverRender = func(c *gin.Context, code int, result *response.Result) {
		c.Status(code)
		utils.Render(c, "error", map[string]interface{}{
			"result": result,
		})
	}

	// run
	app.Run()
}
