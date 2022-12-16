package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/sariyanta/goweb/pkg/config"
	"github.com/sariyanta/goweb/pkg/models"
)

var app *config.AppConfig

func AppDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	//  get the template cache from the app config
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buff := new(bytes.Buffer)

	td = AppDefaultData(td)

	_ = t.Execute(buff, td)

	// render the template
	_, err := buff.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}

	// get all the files named *.page.tmpl
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return templateCache, err
	}

	// loop through the pages
	for _, page := range pages {
		// get the file name
		name := filepath.Base(page)

		// parse the template
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return templateCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return templateCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return templateCache, err
			}
		}

		// add to template cache
		templateCache[name] = ts
	}

	return templateCache, nil
}

func NewTemplates(a *config.AppConfig) {
	app = a
}
