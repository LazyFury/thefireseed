package model

import (
	"strings"

	"github.com/lazyfury/go-web-template/model"
	"github.com/lazyfury/go-web-template/response"
)

type AppModel struct {
	model.Model
	UserID         uint   `json:"user_id" gorm:"not_null"`
	AppName        string `json:"app_name" gorm:"not_null"`
	AppDesc        string `json:"app_desc" gorm:"type:text"`
	AppIcon        string `json:"app_icon"`
	AppAuthor      string `json:"app_author"`
	AppAuthorEmail string `json:"app_author_email"`
}

var _ model.Controller = &AppModel{}

func (a *AppModel) Validator() error {

	a.AppName = strings.Trim(a.AppName, " ")
	if a.AppName == "" {
		response.Error("请输入应用名称")
	}

	return nil
}
func (a *AppModel) Object() interface{} {
	return &AppModel{}
}
func (a *AppModel) Objects() interface{} {
	return &[]AppModel{}
}
func (a *AppModel) TableName() string {
	return "apps"
}
