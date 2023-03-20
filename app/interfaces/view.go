package interfaces

import (
	"kanko-hackaton-22/app/interfaces/handler"

	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func viewRouter(e *echo.Echo) {
	t := &Template{
		templates: template.Must(template.ParseGlob("assets/*.html")),
	}

	e.Renderer = t

	e.GET("/spots", handler.Spots)
}
