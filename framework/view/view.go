package view

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

var resourceTemplatePath string = "resources/templates"

func Render(w http.ResponseWriter, templatePath string) {
	//templatePath := "resources/templates"

	absTemplatePath, err := filepath.Abs(resourceTemplatePath + string(os.PathSeparator) + templatePath)

	if err != nil {
		panic(err)
	}

	t, _ := template.ParseFiles(absTemplatePath)
	t.Execute(w, "")
}
