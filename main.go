package main

import (
	"flag"
	"log"
	"main/database"
	"main/handlers"
	"net/http"
	"path/filepath"

	"github.com/caddyserver/certmagic"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	database.Init()
	defer database.Close()

	var productionmode bool
	flag.BoolVar(&productionmode, "production", false, "production mode")
	flag.Parse()

	handlers.InitTemplates(productionmode)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// assets (css, js, images, etc.)
	router.Group(func(router chi.Router) {
		router.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir(filepath.Join("static", "assets")))))
	})

	// public routes grouped to easily add rate limiting and more
	router.Group(func(router chi.Router) {
		router.Get("/", handlers.IndexPage)
	})

	// private routes requiring authentication
	router.Group(func(router chi.Router) {
		router.Use(middleware.BasicAuth("admin", map[string]string{"admin": "password"}))
		router.Get("/dashboard", handlers.DashboardPage)
	})

	if !productionmode {
		// launching site in debug mode
		log.Println("Starting server http://localhost:1337")
		log.Panic(http.ListenAndServe("localhost:1337", router))
	} else {
		// launching site in production mode with SSL / HTTPS
		domain := "example.com"
		certmagic.DefaultACME.Agreed = true
		certmagic.DefaultACME.Email = "my-email@example.com"
		log.Println("Starting server https://" + domain)
		log.Panic(certmagic.HTTPS([]string{domain, "www." + domain}, router))
	}
}
