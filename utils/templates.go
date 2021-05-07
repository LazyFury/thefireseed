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

	Link struct {
		URL  string
		Name string
	}

	Banner struct {
		Title string
		Tips  string
	}
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
func (s *SEO) SetKeywords(keywords string) *SEO {
	s.Keywords = keywords
	return s
}
func (s *SEO) SetDesc(desc string) *SEO {
	s.Desc = desc
	return s
}

func parseRenderParams(args ...interface{}) (data UserParam, seo *SEO, banner *Banner) {
	data = UserParam{}
	for _, arg := range args {
		// data 里的数据
		// log.Print(reflect.ValueOf(arg).Type())
		dataArr, ok := arg.(UserParam)
		if ok {
			for k, v := range dataArr {
				(data)[k] = v
			}
			continue
		}
		pointerdataArr, ok := arg.(*UserParam)
		if ok {
			for k, v := range *pointerdataArr {
				(data)[k] = v
			}
			continue
		}
		// seo
		_seo, ok := arg.(*SEO)
		if ok {
			seo = _seo
			continue
		}

		_banner, ok := arg.(Banner)
		if ok {
			banner = &_banner
			continue
		}

		_pointerbanner, ok := arg.(*Banner)
		if ok {
			banner = _pointerbanner
			continue
		}
	}
	return
}

var (
	// 模板集 html/template 扫描文件并加载到内存到模版集合
	Tmpl = template.Must(_template.ParseGlob(template.New("main").Funcs(TemplateFuncs), "./templates", "*.html"))
	// 一个自定义到layout 搭配通用页头页脚，通用的参数
	_layout = layout.New("home/base.html", Tmpl)
	nav     = []Link{
		{"/", "首页"},
		{"/fireseed", "火种🔥"},
		{"/about", "关于我们"},
	}
	links = []Link{
		{"/", "indiaDev"},
		{"/", "v2ex"},
		{"/", "juejin"},
		{"/", "fish"},
		{"/", "debian"},
	}
)

// Render 渲染模版并提供公共参数
// @params args   SEO,UserParam
func Render(c *gin.Context, name string, args ...interface{}) {
	data, seo, banner := parseRenderParams(args...)

	user := controller.GetUserOrEmpty(c)

	// seo
	if seo == nil {
		seo = DefaultSEO()
	}

	params := &layout.LayoutParams{
		TemplateName: name,
		Data:         data,
		Header: map[string]interface{}{
			"user":   user,
			"seo":    seo,
			"nav":    nav,
			"logo":   "🔥火种计划",
			"banner": banner,
		},
		Footer: map[string]interface{}{
			"links": links,
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
