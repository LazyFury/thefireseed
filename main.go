package main

import (
	"fmt"
	"net/http"
	"thefireseed/config"
	"thefireseed/middleware"
	"thefireseed/model"
	"thefireseed/router"
	"thefireseed/utils"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/kr/pretty"
	gwt "github.com/lazyfury/go-web-template"
	"github.com/lazyfury/go-web-template/response"
)

func main() {
	// init
	app := gwt.New()
	pretty.Print(map[string]interface{}{"test": "pretty print"})
	// 连接数据库
	if err := model.DB.ConnectMysql(config.Global.Mysql.ToString()); err != nil {
		panic(err)
	}

	model.DB.AutoMigrate(
		&model.AppModel{}, &model.ActivityModel{}, &model.CodeModel{},
		&model.User{}, &model.CodeCopyLogModel{},
	)
	app.PreUse(func(c *gin.Context) {
		// pretty.Print(c.Request)
	})

	app.PreUse(middleware.AuthOrNot, gwt.DefaultCors)

	// 注册模版
	app.SetHTMLTemplate(utils.Tmpl)

	// 注册静态目录
	app.Use(static.Serve("/static", static.LocalFile("static", false)))

	// 注册路由
	app.InitRouter(router.Init)

	// 没有注册的路由，不走全局中间件，TODO:
	app.NoMethodUse(middleware.AuthOrNot)
	app.NoRouteUse(middleware.AuthOrNot)
	// 错误码配置
	response.RecoverErrHtml = true
	response.RecoverRender = func(c *gin.Context, code int, result *response.Result) {
		c.Status(http.StatusOK)
		utils.Render(c, "error", utils.UserParam{
			"result": result,
		})
	}
	// run
	app.Run(fmt.Sprintf(":%d", config.Global.Prot))
}
