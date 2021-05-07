package model

import (
	"net/http"
	"strings"

	"github.com/lazyfury/go-web-template/model"
	"github.com/lazyfury/go-web-template/response"
)

type CodeUsedLogModel struct {
	model.Model
	InviteId string `json:"invite_code" gorm:"conment:邀请码的id;"` //邀请码的id
	UserCode string `json:"user_code"  gorm:"conment:用户id;"`    //用户id
	Data     string `json:"data" gorm:"conment:邀请码;"`           //邀请码
}

var _ model.Controller = &CodeUsedLogModel{}

func (a *CodeUsedLogModel) Validator() error {
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

	copy := &CodeUsedLogModel{}
	if err := DB.GetObjectOrNotFound(copy, map[string]interface{}{
		"invite_id": invite.Code,
		"user_code": a.UserCode,
	}); err == nil {
		response.Error(response.JSON(http.StatusAlreadyReported, "已经复制过了", copy))
	}

	return nil
}
func (a *CodeUsedLogModel) Object() interface{} {
	return &CodeUsedLogModel{}
}
func (a *CodeUsedLogModel) Objects() interface{} {
	return &[]CodeUsedLogModel{}
}
func (a *CodeUsedLogModel) TableName() string {
	return "code_used_logs"
}
