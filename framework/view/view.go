package view

import (
	//	"github.com/davecgh/go-spew/spew"
	"html/template"
	"net/http"
	"path/filepath"
)

var resourceTemplatePath string = "resources/templates"

type View struct {
	templates *template.Template
}

func NewView() View {
	v := &View{}
	v.LoadViews()

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

func (v *View) Render(w http.ResponseWriter, r *http.Request, templateName string, data interface{}) {
	err := v.templates.ExecuteTemplate(w, templateName, data)

	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
	}
}
