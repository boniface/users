package login

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"users/cmd/config"
)

// Rout Path
func Login(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", LoginHandler(app))
	r.Post("/accounts", GetAccountsHandler(app))
	return r

}

func LoginHandler(app *config.Env) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			"./views/html/login.page.html",
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

func Logout(app *config.Env) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("welcome"))

	}
}

func ForgotPassword(app *config.Env) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("welcome"))

	}
}

func GetAccountsHandler(app *config.Env) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("welcome"))

	}
}
