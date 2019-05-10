package security

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"users/cmd/config"
)

func Security(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", SecurityHanler(app))
	return r

}

func SecurityHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			"./views/html/security/security.page.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}
