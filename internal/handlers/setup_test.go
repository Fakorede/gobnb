package handlers

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/justinas/nosurf"

	"github.com/alexedwards/scs/v2"
	"github.com/fakorede/gobnb/internal/config"
	"github.com/fakorede/gobnb/internal/models"
	"github.com/fakorede/gobnb/internal/render"
)

var session *scs.SessionManager
var app config.AppConfig

var functions = template.FuncMap{}
var pathToTemplates = "./../../templates"

func getRoutes() http.Handler {
	// values to be stored in session
	gob.Register(models.Reservation{})

	app.InProduction = false

	// configure logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// Initialize a new session manager and configure the session lifetime.
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = true

	repo := NewRepo(&app)

	NewHandlers(repo)
	render.NewRenderer(&app)

	mux := chi.NewRouter()

	// mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	// mux.Use(CsrfTokenMiddleware)
	mux.Use(SessionMiddleware)

	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)
	mux.Get("/general-quarters", Repo.Generals)
	mux.Get("/major-suites", Repo.Majors)
	mux.Get("/contact", Repo.Contact)

	mux.Get("/make-reservation", Repo.Reservation)
	mux.Post("/make-reservation", Repo.MakeReservation)
	mux.Get("/reservation-summary", Repo.ReservationSummary)

	mux.Get("/search-availability", Repo.Availability)
	mux.Post("/search-availability-json", Repo.CheckAvailabilityJSON)
	mux.Post("/search-availability", Repo.CheckAvailability)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux

}

// CsrfTokenMiddleware adds CSRF protection to all post requests
func CsrfTokenMiddleware(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionMiddleware loads and saves the session on every request
func SessionMiddleware(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

// CreateTestTemplateCache creates a template cache as a map
func CreateTestTemplateCache() (map[string]*template.Template, error) {
	tempCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return tempCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		tempSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return tempCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			return tempCache, err
		}

		if len(matches) > 0 {
			tempSet, err = tempSet.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return tempCache, err
			}
		}

		tempCache[name] = tempSet // fully parsed ready-to-use template
	}

	return tempCache, nil
}
