package view

import (
	//	"github.com/davecgh/go-spew/spew"
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

var resourceTemplatePath string = "resources/templates"

var layoutFuncs = template.FuncMap{
	"yield": func() (string, error) {
		return "", fmt.Errorf("yield called inappropriately")
	},
}

type View struct {
	templates   *template.Template
	adminLayout *template.Template
}

func NewView() View {
	v := &View{}
	v.LoadViews()
	v.LoadLayouts()
	return *v
}

func (v *View) LoadViews() {
	absTemplatePath, err := filepath.Abs(resourceTemplatePath + "/**/**/*.html")

	if err != nil {
		panic(err)
	}

	var templates = template.Must(template.New("t").ParseGlob(absTemplatePath))

	v.templates = templates
}

func (v *View) LoadLayouts() {
	v.adminLayout = template.Must(
		template.
			New("layout.html").
			Funcs(layoutFuncs).
			ParseFiles(resourceTemplatePath + "/admin/layout.html"),
	)
}

func (v *View) Render(w http.ResponseWriter, r *http.Request, templateName string, data interface{}) {

	if strings.Index(templateName, "admin") == 0 {
		v.RenderAdmin(w, r, templateName, data)
	}

}

func (v *View) RenderAdmin(w http.ResponseWriter, r *http.Request, templateName string, data interface{}) {

	// Override the yield func so that we can inject the partial template
	funcs := template.FuncMap{
		"yield": func() (template.HTML, error) {
			buff := bytes.NewBuffer(nil)
			err := v.templates.ExecuteTemplate(buff, templateName, data)
			return template.HTML(buff.String()), err
		},
	}

	adminLayoutClone, _ := v.adminLayout.Clone()
	adminLayoutClone.Funcs(funcs)
	err := adminLayoutClone.Execute(w, data)

	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
	}
}
