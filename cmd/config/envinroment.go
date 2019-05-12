package config

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"runtime/debug"
	"users/pkg/forms"
)

type Env struct {
	ErrorLog      *log.Logger
	InfoLog       *log.Logger
	TemplateCache map[string]*template.Template
}

type templateData struct {
	Title string
	Data  map[string]string
	Form  *forms.Form
}

func (app *Env) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// For consistency, we'll also implement a notFound helper. This is simply a
// convenience wrapper around clientError which sends a 404 Not Found response to
// the user.
func (app *Env) NotFound(w http.ResponseWriter) {
	app.ClientError(w, http.StatusNotFound)
}

func (app *Env) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Env) AddDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}

	return td
}

func (app *Env) Render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.TemplateCache[name]
	if !ok {
		app.ServerError(w, fmt.Errorf("The template %s does not exist", name))
		return
	}

	buf := new(bytes.Buffer)

	err := ts.Execute(buf, app.AddDefaultData(td, r))
	if err != nil {
		app.ServerError(w, err)
		return
	}

	buf.WriteTo(w)
}

func NewTemplateCache(dir string) (map[string]*template.Template, error) {
	// Initialize a new map to act as the cache.
	cache := map[string]*template.Template{}

	// Use the filepath.Glob function to get a slice of all filepaths with
	// the extension '.page.tmpl'. This essentially gives us a slice of all the
	// 'page' templates for the application.
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.html"))
	if err != nil {
		return nil, err
	}

	// Loop through the pages one-by-one.
	for _, page := range pages {
		// Extract the file name (like 'home.page.tmpl') from the full file path
		// and assign it to the name variable.
		name := filepath.Base(page)

		// Parse the page template file in to a template set.
		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Use the ParseGlob method to add any 'layout' templates to the
		// template set (in our case, it's just the 'base' layout at the
		// moment).
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.html"))
		if err != nil {
			return nil, err
		}

		// Use the ParseGlob method to add any 'partial' templates to the
		// template set (in our case, it's just the 'footer' partial at the
		// moment).
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.html"))
		if err != nil {
			return nil, err
		}

		// Add the template set to the cache, using the name of the page
		// (like 'home.page.tmpl') as the key.
		cache[name] = ts
	}

	// Return the map.
	return cache, nil
}
