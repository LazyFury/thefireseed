package utils

import (
	"fmt"
	"html/template"

	"github.com/lazyfury/go-web-template/tools"
)

type (
	// 分页Item
	PageItem struct {
		IsCurrent bool
		Text      string
		URL       string
	}
)

var (
	Bootstrap = template.Must(tools.ParseGlob(template.New("main").Funcs(TemplateFuncs), "templates", "*.html"))
)

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
}
