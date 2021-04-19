package model

import (
	"strings"

	"github.com/lazyfury/go-web-template/model"
	"github.com/lazyfury/go-web-template/response"
)

type CodeModel struct {
	model.Model
	ActivityCode uint   `json:"activety_code" gorm:"not_null"`
	InviteCode   string `json:"code" gorm:"not_null"`
	Used         bool   `json:"is_used"`
}

var _ model.Controller = &CodeModel{}

func (a *CodeModel) Validator() error {
	a.Code = strings.Trim(a.Code, " ")
	if a.Code == "" {
		response.Error("请输入活动名称")
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
