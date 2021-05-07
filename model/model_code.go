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
	}
	ShowCodeModal struct {
		*CodeModel
		Count     int  `json:"copy_count" gorm:"column:copy_count;"`
		UsedCount int  `json:"used_count" gorm:"column:used_count;"`
		Copied    bool `json:"copied"`
		Used      bool `json:"used"`
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
	list, ok := data.(*[]ShowCodeModal)
	if ok {
		for i, invite := range *list {
			// hidden code
			str := invite.InviteCode
			strLen := len(str)
			if strLen >= 8 {
				b := []byte{}
				for i := 0; i < (strLen - 6); i++ {
					b = append(b, []byte("*")...)
				}
				str = str[:3] + string(b) + str[strLen-3:]
			}
			(*list)[i].InviteCode = str
		}
	}
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
	used := &CodeUsedLogModel{}
	db = db.Joins(fmt.Sprintf("left join (select count(id) `copy_count`,`invite_id` `t1_invite_code` from `%s` group by `invite_id`) t1 on t1.`t1_invite_code`=`%s`.`code`", copy.TableName(), a.TableName()))
	db = db.Joins(fmt.Sprintf("left join (select count(id) `used_count`,`invite_id` `t2_invite_code` from `%s` group by `invite_id`) t2 on t2.`t2_invite_code`=`%s`.`code`", used.TableName(), a.TableName()))
	return db
}

func (a *CodeModel) TableName() string {
	return "codes"
}
