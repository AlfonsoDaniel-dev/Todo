package config

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

// TODO implement templates configuration

type templates struct {
	Templates *template.Template
	PagesDir  string
}

func (t *templates) Render(w io.Writer, templateName string, templateData interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, templateName, templateData)
}

func NewTemplates(pagesDir string) *templates {

	t := &templates{
		PagesDir: pagesDir,
	}

	directory := pagesDir + "/*.html"
	t.Templates = template.Must(template.ParseGlob(directory))

	return t
}
