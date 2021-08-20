package render

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/fakorede/gobnb/internal/config"
	"github.com/fakorede/gobnb/internal/models"

	"github.com/justinas/nosurf"
)

// map of functions we can use in a template
var functions = template.FuncMap{}

var app *config.AppConfig
var pathToTemplates = "./templates"

// sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

// RenderTemplate renders template using the html/template package
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) error {
	var tc map[string]*template.Template

	if app.UseCache {
		// get template cache from app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}


	t, ok := tc[tmpl]
	if !ok {
		// log.Fatal("Could not get template from template cache")
		return errors.New("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	err := t.Execute(buf, td)
	if err != nil {
		log.Fatal(err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser:", err)
		return err
	}

	return nil
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
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