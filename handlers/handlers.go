package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
	"time"
)

var templates *template.Template

func InitTemplates(productionmode bool) {
	go func() {
		for {
			templates = template.Must(template.ParseGlob(filepath.Join("static", "html", "pages", "*.html")))
			templates = template.Must(templates.ParseGlob(filepath.Join("static", "html", "components", "*.html")))

			if !productionmode {
				time.Sleep(3 * time.Second) // reload templates every 3 seconds in debug mode
			} else {
				time.Sleep(12 * time.Hour) // reload templates every 12 hours in production mode
			}
		}
	}()
}

func IndexPage(res http.ResponseWriter, req *http.Request) {
	templates.ExecuteTemplate(res, "index.html", nil)
}

func DashboardPage(res http.ResponseWriter, req *http.Request) {
	templates.ExecuteTemplate(res, "dashboard.html", nil)
}
