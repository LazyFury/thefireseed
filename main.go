package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"thefireseed/config"
	"thefireseed/middleware"
	"thefireseed/model"
	"thefireseed/router"
	"thefireseed/utils"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	gwt "github.com/lazyfury/go-web-template"
	"github.com/lazyfury/go-web-template/response"
)

func redirectLog() *os.File {
	f, _ := os.Create("_.log")
	os.Stdout = f
	os.Stderr = f
	log.SetOutput(f)
	log.SetFlags(log.Lshortfile)
	return f
}
func main() {
	f := redirectLog()
	defer f.Close()
	// init
	app := gwt.New()
	// pretty.Println(map[string]interface{}{"test": "pretty print"})
	// 连接数据库
	if err := model.DB.ConnectMysql(config.Global.Mysql.ToString()); err != nil {
		panic(err)
	}

	model.DB.AutoMigrate(
		&model.AppModel{}, &model.ActivityModel{}, &model.CodeModel{},
		&model.User{},
		&model.CodeCopyLogModel{}, &model.CodeUsedLogModel{},
	)

	// log
	app.PreUse(gin.LoggerWithConfig(gin.LoggerConfig{
		Output: f,
	}))

	// cors
	app.PreUse(func(c *gin.Context) {
		gwt.Cors(c, &config.Application.CORS)
		c.Next()
	})

	// auth
	app.PreUse(middleware.AuthOrNot)

	// 注册模版
	app.SetHTMLTemplate(utils.Tmpl)

	// 注册静态目录
	app.Use(static.Serve("/static", static.LocalFile("static", false)))

	// 注册路由
	app.InitRouter(router.Init)

	// 没有注册的路由，不走全局中间件
	app.NoMethodUse(middleware.AuthOrNot)
	app.NoRouteUse(middleware.AuthOrNot)
	// 根据content-type为text/html时返回页面而不是json
	response.RecoverErrHtml = true
	response.RecoverRender = func(c *gin.Context, code int, result *response.Result) {
		c.Status(http.StatusOK)
		utils.Render(c, "error", utils.UserParam{
			"result": result,
		})
	}
	// run
	app.Run(fmt.Sprintf(":%d", config.Application.Prot))
}
