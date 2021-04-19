package utils

import (
	"fmt"
	"shareInviteCode/controller"
	"shareInviteCode/model"

	"github.com/gin-gonic/gin"
	"github.com/lazyfury/go-web-template/tools/template/layout"
)

type (
	// 分页Item
	PageItem struct {
		IsCurrent bool
		Text      string
		URL       string
	}
)

func Render(c *gin.Context, name string, args ...map[string]interface{}) {
	layout := layout.New(name, args...)

	user := controller.GetUserOrEmpty(c)
	layout.Header["user"] = user

	layout.Render(c)
}

var TemplateFuncs = map[string]interface{}{
	"plus": func(x int, y int) int {
		return x + y
	},
	"reduce": func(x int, y int) int {
		return x - y
	},
	"StrJoin": func(str string, args ...interface{}) string {
		return fmt.Sprintf(str, args...)
	},
	"NewPageItem": func(isCurrent bool, text string, url string) *PageItem {
		return &PageItem{
			IsCurrent: isCurrent,
			Text:      text,
			URL:       url,
		}
	},
	"hasUser": func(user *model.User) bool {
		return user.ID > 0
	},
}
