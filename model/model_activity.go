package model

import (
	"strings"

	"github.com/lazyfury/go-web-template/model"
	"github.com/lazyfury/go-web-template/response"
)

type ActivityModel struct {
	model.Model
	AppCode      uint   `json:"app_code"`
	ActivityName string `json:"activity_name"`
}

var _ model.Controller = &ActivityModel{}

func (a *ActivityModel) Validator() error {

	a.ActivityName = strings.Trim(a.ActivityName, " ")
	if a.ActivityName == "" {
		response.Error("请输入活动名称")
	}

	return nil
}
func (a *ActivityModel) Object() interface{} {
	return &ActivityModel{}
}
func (a *ActivityModel) Objects() interface{} {
	return &[]ActivityModel{}
}
func (a *ActivityModel) TableName() string {
	return "activitys"
}
