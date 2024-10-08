package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/sagar2395/golang-by-trevor-sawler/pkg/config"
	"github.com/sagar2395/golang-by-trevor-sawler/pkg/models"
)

func RenderTemplatesTest(w http.ResponseWriter, tmpl string) {
	fmt.Println("inside render template function")
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("Error Parsing template", err)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplates(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if we already have a template
	_, inMap := tc[t]

	if !inMap {

		log.Println("Creating template and adding to cache")
		err = CreateTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
		// need to create the template
	} else {
		// We have template already in cache
		log.Println("using cached template")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}

	tc[t] = tmpl

	return nil
}

// Approach 2 for Rendering templates by Caching items.

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	// Add all the data which should be added by default to the page

	return td
}

func RenderTemplatesApproach2(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// Create a template Cache
	// tc, err := CreateTemplateCacheApproach2()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCacheApproach2()
	}

	// Get requested template from cache
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("Could not get template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	//Render the template
	fmt.Println("inside render template function")
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCacheApproach2() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// Get all files named *.page.tmpl from ./templates

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	//range through all files ending with *.page.tmpl

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
