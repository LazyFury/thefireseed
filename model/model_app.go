package model

import (
	"github.com/lazyfury/go-web-template/model"
)

type AppModel struct {
	model.Model
	Name string `json:"app_name"`
}

var _ model.Controller = &AppModel{}

func (a *AppModel) Validator() error {
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
