package model

import (
	"strings"

	"github.com/lazyfury/go-web-template/model"
	"github.com/lazyfury/go-web-template/response"
)

type CodeCopyLogModel struct {
	model.Model
	InviteId string `json:"invite_code" gorm:"conment:邀请码的id;"` //邀请码的id
	UserCode string `json:"user_code"  gorm:"conment:用户id;"`    //用户id
	Data     string `json:"data" gorm:"conment:邀请码;"`           //邀请码
}

var _ model.Controller = &CodeCopyLogModel{}

func (a *CodeCopyLogModel) Validator() error {
	a.InviteId = strings.Trim(a.InviteId, " ")
	if a.InviteId == "" {
		response.Error("邀请码code不可空")
	}

	invite := &CodeModel{}
	if err := DB.GetObjectOrNotFound(invite, map[string]interface{}{
		"code": a.InviteId,
	}); err != nil {
		response.Error(err)
	}

	a.Data = invite.InviteCode
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
