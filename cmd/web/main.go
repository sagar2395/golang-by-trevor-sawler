package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/sagar2395/golang-by-trevor-sawler/pkg/config"
	"github.com/sagar2395/golang-by-trevor-sawler/pkg/handlers"
	"github.com/sagar2395/golang-by-trevor-sawler/pkg/render"
)

const portNumber = ":8090"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	var app config.AppConfig

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCacheApproach2()

	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	if err != nil {
		log.Fatal("Can not create template cache")
	}
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	srv.ListenAndServe()
	log.Fatal(err)

	// _ = http.ListenAndServe(":8090", nil)
}
