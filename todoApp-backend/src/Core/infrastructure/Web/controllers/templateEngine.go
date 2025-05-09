package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"log"
	"path/filepath"
)

type templates struct {
	Templates *template.Template
	PagesDir  string
}

func (t *templates) Render(w io.Writer, templateName string, templateData interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, templateName, templateData)
}

func newTemplates(pagesDir string) *templates {

	t := &templates{
		PagesDir: pagesDir,
	}

	directory := filepath.Join(pagesDir, "*.html")
	fmt.Println(directory)
	templatesInstance, err := template.ParseGlob(directory)
	if err != nil {
		log.Fatalf("Error parsing templates: %s", err)
	}

	t.Templates = templatesInstance

	return t
}
