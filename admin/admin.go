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

		if err := a.db.FindAll(model); err != nil {
			ctx.HTML(http.StatusInternalServerError, "error.html", nil)
			return
		}

		ctx.HTML(http.StatusAccepted, "list.html", map[string]interface{}{"items": model, "name": name})
	}
}

func (a *Admin) newHandler(name string) HandlerFunc {
	return func(ctx Context) {
		ctx.HTML(http.StatusOK, "form.html", map[string]interface{}{"action": "admin/" + name, "method": "POST"})
	}
}

func (a *Admin) createHandler(name string, model interface{}) HandlerFunc {
	return func(ctx Context) {
		if err := ctx.Bind(model); err == nil {
			if err := a.db.Create(model); err != nil {
				ctx.HTML(http.StatusInternalServerError, "error.html", map[string]interface{}{"message": "An error occured", "error": err})
				return
			}
			ctx.Redirect(http.StatusFound, "/admin/"+name)
		} else {
			ctx.HTML(http.StatusInternalServerError, "error.html", map[string]interface{}{"message": "An error occured", "error": err})
		}
	}
}
