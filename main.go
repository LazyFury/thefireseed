package main

import (
	"html/template"
	"shareInviteCode/config"
	"shareInviteCode/model"
	"shareInviteCode/router"

	"github.com/gin-contrib/static"
	gowebtemplate "github.com/lazyfury/go-web-template"
	"github.com/lazyfury/go-web-template/response"
	"github.com/lazyfury/go-web-template/tools"
)

func main() {
	// init
	app := gowebtemplate.New()

	// 连接数据库
	if err := model.DB.ConnectMysql(config.Global.Mysql.ToString()); err != nil {
		panic(err)
	}

	model.DB.AutoMigrate(
		&model.AppModel{}, &model.ActivityModel{}, &model.CodeModel{},
		&model.User{},
	)

	// 注册模版
	html := template.Must(tools.ParseGlob(template.New("main"), "templates", "*.html"))
	app.SetHTMLTemplate(html)

	// 注册静态目录
	app.Use(static.Serve("/", static.LocalFile("wwwroot", false)))

	// 注册路由
	router.Init(&app.RouterGroup)

	// 错误码配置
	response.RecoverErrHtml = false

	// run
	app.Run()
}
