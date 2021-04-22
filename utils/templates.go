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

	SEO struct {
		Title     string
		BaseTitle string
		Keywords  string
		Desc      string
		Author    string
	}

	UserParam map[string]interface{}
)

func DefaultSEO() *SEO {
	return &SEO{
		BaseTitle: "分享邀请码",
		Keywords:  "分享",
		Desc:      "邀请码",
	}
}

func (s *SEO) SetTitle(title string) *SEO {
	s.Title = title
	return s
}

func parseRenderParams(args ...interface{}) (maps []map[string]interface{}, seo *SEO) {
	for _, arg := range args {
		// data 里的数据
		// log.Print(reflect.ValueOf(arg).Type())
		data, ok := arg.(UserParam)
		if ok {
			maps = append(maps, data)
			continue
		}

		// seo
		_seo, ok := arg.(*SEO)
		if ok {
			seo = _seo
			continue
		}

	}
	return
}

// Render 渲染模版并提供公共参数
// @params args   SEO,UserParam
func Render(c *gin.Context, name string, args ...interface{}) {
	maps, seo := parseRenderParams(args...)
	layout := layout.New(name, maps...)

	user := controller.GetUserOrEmpty(c)

	// user
	layout.Header["user"] = user

	// seo
	if seo == nil {
		seo = DefaultSEO()
	}
	layout.Header["seo"] = seo

	// rander
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
	"emptyStr": func(str string) bool {
		return str == ""
	},
}
