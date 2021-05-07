package config

import (
	gwt "github.com/lazyfury/go-web-template"
	"github.com/lazyfury/go-web-template/config"
	"github.com/lazyfury/go-web-template/tools"
	"github.com/lazyfury/go-web-template/tools/mail"
	"github.com/lazyfury/go-web-template/tools/mysql"
)

// Global 全局配置
var Global *baseConfig = config.ReadConfig(&baseConfig{
	Screct: tools.RandStringBytes(32),
}, "./config.json").(*baseConfig)

// Application 配置 TODO:mysql mail port 啥的
var Application *ApplicationConfig = config.ReadConfig(&ApplicationConfig{
	CORS: *gwt.DefaultCorsConfig(),
	Prot: 8080,
}, "./application.json").(*ApplicationConfig)

type ApplicationConfig struct {
	CORS gwt.CorsConfig `json:"cors"`
	Prot int            `json:"port"`
}

type baseConfig struct {
	Mysql  mysql.Mysql `json:"mysql"` // 数据库链接
	Mail   mail.Mail   `json:"mail"`
	Screct string      `json:"screct"`
}
