package utils

import (
	"fmt"
	"html/template"
	"thefireseed/controller"
	"thefireseed/model"

	"github.com/gin-gonic/gin"
	_template "github.com/lazyfury/go-web-template/tools/template"
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

func parseRenderParams(args ...interface{}) (data UserParam, seo *SEO) {
	data = UserParam{}
	for _, arg := range args {
		// data 里的数据
		// log.Print(reflect.ValueOf(arg).Type())
		dataArr, ok := arg.(UserParam)
		if ok {
			for k, v := range dataArr {
				data[k] = v
			}
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

var Tmpl = template.Must(_template.ParseGlob(template.New("main").Funcs(TemplateFuncs), "./templates", "*.html"))
var _layout = layout.New("home/base.html", Tmpl)

// Render 渲染模版并提供公共参数
// @params args   SEO,UserParam
func Render(c *gin.Context, name string, args ...interface{}) {
	data, seo := parseRenderParams(args...)

	user := controller.GetUserOrEmpty(c)

	// seo
	if seo == nil {
		seo = DefaultSEO()
	}

	params := &layout.LayoutParams{
		TemplateName: name,
		Data:         data,
		Header: map[string]interface{}{
			"user": user,
			"seo":  seo,
		},
	}

	_layout.Render(c, params)
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
