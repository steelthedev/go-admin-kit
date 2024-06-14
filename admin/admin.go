package admin

import (
	"embed"
	"html/template"
	"net/http"
)

// go:embed templates/*.html
var content embed.FS

type Admin struct {
	db        Database
	router    Router
	templates *template.Template
}

func NewAdmin(db Database, router Router, templates *template.Template) *Admin {
	templ := template.Must(template.New("").ParseFS(content, "templates/*.html"))
	return &Admin{
		db,
		router,
		templ,
	}
}

func (a *Admin) RegisterModel(name string, model interface{}) {
	a.setUproutes(name, model)
}

func (a *Admin) setUproutes(name string, model interface{}) {
	a.router.GET("/admin/"+name, a.listHandler(name, model))
}

func (a *Admin) listHandler(name string, model interface{}) HandlerFunc {
	return func(ctx Context) {
		var results []interface{}

		if err := a.db.FindAll(results); err != nil {
			ctx.HTML(http.StatusInternalServerError, "error.html", nil)
			return
		}

		ctx.HTML(http.StatusAccepted, "list.html", map[string]interface{}{"items": results, "name": name})
	}
}
