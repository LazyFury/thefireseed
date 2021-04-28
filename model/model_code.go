package model

import (
	"fmt"
	"strings"

	"github.com/lazyfury/go-web-template/model"
	"github.com/lazyfury/go-web-template/response"
	"gorm.io/gorm"
)

type (
	CodeModel struct {
		model.Model
		ActivityCode string `json:"activity_code" gorm:"not_null"`
		InviteCode   string `json:"code" gorm:"not_null"`
		Used         bool   `json:"is_used"`
	}
	ShowCodeModal struct {
		*CodeModel
		Count  int  `json:"count"`
		Copied bool `json:"copied"`
	}
)

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

func (a *CodeModel) Result(data interface{}) interface{} {
	return data
}

func (a *CodeModel) Object() interface{} {
	return &CodeModel{}
}
func (a *CodeModel) Objects() interface{} {
	return &[]ShowCodeModal{}
}

func (a *CodeModel) Joins(db *gorm.DB) *gorm.DB {
	copy := &CodeCopyLogModel{}
	db = db.Joins(fmt.Sprintf("left join (select count(id) `count`,`invite_id` `t1_invite_code` from `%s` group by `invite_id`) t1 on t1.`t1_invite_code`=`%s`.`code`", copy.TableName(), a.TableName()))
	return db
}

func (a *CodeModel) TableName() string {
	return "codes"
}
