package main

import (
	"shareInviteCode/config"
	"shareInviteCode/model"
	"shareInviteCode/router"
	"shareInviteCode/utils"

	"github.com/gin-contrib/static"
	gwt "github.com/lazyfury/go-web-template"
	"github.com/lazyfury/go-web-template/response"
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
	app.SetHTMLTemplate(utils.Bootstrap)

	// 注册静态目录
	app.Use(static.Serve("/static", static.LocalFile("static", false)))

	// 注册路由
	router.Init(&app.RouterGroup)

	// 错误码配置
	response.RecoverErrHtml = false
	response.RecoverErrTemplateName = "err/error.html"
	// run
	app.Run()
}
