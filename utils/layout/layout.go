package layout

import (
	"bytes"
	"html/template"
	"net/http"
	"shareInviteCode/utils"

	"github.com/gin-gonic/gin"
	"github.com/lazyfury/go-web-template/response"
)

type (
	LayoutParams struct {
		Header             map[string]interface{}
		Data               map[string]interface{}
		Footer             map[string]interface{}
		TemplateName       string
		HeaderTemplateName string
		FooterTemplateName string
		Template           *template.Template
	}
)

func New(name string, args ...map[string]interface{}) *LayoutParams {
	data := map[string]interface{}{}

	for _, arg := range args {
		for k := range arg {
			data[k] = arg[k]
		}
	}

	return &LayoutParams{
		Data:         data,
		TemplateName: name,
		Template:     utils.Bootstrap,
	}
}
func Render(c *gin.Context, name string, args ...map[string]interface{}) {
	layout := New(name, args...)
	layout.Render(c)
}
func (p *LayoutParams) Render(c *gin.Context) {

	w := bytes.NewBuffer([]byte(""))
	if p.HeaderTemplateName == "" {
		p.HeaderTemplateName = "header"
	}
	if p.FooterTemplateName == "" {
		p.FooterTemplateName = "footer"
	}
	tmpl := p.Template.New("layout")
	err := tmpl.ExecuteTemplate(w, p.HeaderTemplateName, p.Header)
	if err != nil {
		response.Error(err)
	}
	err = tmpl.ExecuteTemplate(w, p.TemplateName, p.Data)
	if err != nil {
		response.Error(err)
	}
	err = tmpl.ExecuteTemplate(w, p.FooterTemplateName, p.Footer)
	if err != nil {
		response.Error(err)
	}
	html := template.HTML(w.String())
	c.Status(http.StatusOK)
	_, err = c.Writer.Write([]byte(html))
	if err != nil {
		response.Error(err)
	}
}
