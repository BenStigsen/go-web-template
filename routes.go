package main

import "net/http"

func IndexPage(res http.ResponseWriter, req *http.Request) {
	templates.ExecuteTemplate(res, "index.html", nil)
}

func DashboardPage(res http.ResponseWriter, req *http.Request) {
	templates.ExecuteTemplate(res, "dashboard.html", nil)
}
