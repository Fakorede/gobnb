package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fakorede/gobnb/pkg/config"
	"github.com/fakorede/gobnb/pkg/handlers"
	"github.com/fakorede/gobnb/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on Port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
