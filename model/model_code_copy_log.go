package model

import (
	"strings"

	"github.com/lazyfury/go-web-template/model"
	"github.com/lazyfury/go-web-template/response"
)

type CodeCopyLogModel struct {
	model.Model
	InviteCode string `json:"invite_code"`
	UserCode   string `json:"user_code"`
}

var _ model.Controller = &CodeCopyLogModel{}

func (a *CodeCopyLogModel) Validator() error {
	a.InviteCode = strings.Trim(a.InviteCode, " ")
	if a.InviteCode == "" {
		response.Error("邀请码code不可空")
	}
	a.UserCode = strings.Trim(a.UserCode, " ")
	if a.UserCode == "" {
		response.Error("用户code不可空")
	}

	return nil
}
func (a *CodeCopyLogModel) Object() interface{} {
	return &CodeCopyLogModel{}
}
func (a *CodeCopyLogModel) Objects() interface{} {
	return &[]CodeCopyLogModel{}
}
func (a *CodeCopyLogModel) TableName() string {
	return "code_copy_logs"
}
