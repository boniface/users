package login

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"users/cmd/config"
	"users/cmd/web/middleware"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequireAuthenticatedUser)
	r.Get("/", IndexHanler(app))
	return r

}

func IndexHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			"./views/html/index.html",
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
