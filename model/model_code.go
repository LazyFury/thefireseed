package model

import (
	"strings"

	"github.com/lazyfury/go-web-template/model"
	"github.com/lazyfury/go-web-template/response"
)

type CodeModel struct {
	model.Model
	ActivityCode string `json:"activity_code" gorm:"not_null"`
	InviteCode   string `json:"code" gorm:"not_null"`
	Used         bool   `json:"is_used"`
}

var _ model.Controller = &CodeModel{}

func (a *CodeModel) Validator() error {

	a.ActivityCode = strings.Trim(a.ActivityCode, " ")
	if a.ActivityCode == "" {
		response.Error("请输入活动Code")
	}

	a.InviteCode = strings.Trim(a.InviteCode, " ")
	if a.InviteCode == "" {
		response.Error("请输入激活码")
	}

	return nil
}
func (a *CodeModel) Object() interface{} {
	return &CodeModel{}
}
func (a *CodeModel) Objects() interface{} {
	return &[]CodeModel{}
}
func (a *CodeModel) TableName() string {
	return "codes"
}
