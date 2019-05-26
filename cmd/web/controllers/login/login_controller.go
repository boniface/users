package login

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"users/cmd/config"
	"users/pkg/api/login"
)

// Route Path
func Login(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", LoginHandler(app))
	r.Post("/accounts", GetAccountsHandler(app))
	r.Get("/password", PasswordHandler(app))
	return r

}

func LoginHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "/login/login.page.html",
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

func PasswordHandler(app *config.Env) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		files := []string{
			app.Path + "/login/password.page.html",
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

func GetAccountsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var files []string
		r.ParseForm()
		email := r.PostFormValue("email")
		accounts, err := login.GetUserAccounts(email)
		if err != nil {
			fmt.Println(" The Error ", err)
		}
		switch 2 {
		case 0:
			{
				http.Redirect(w, r, "/login", 301)
			}
		case 1:
			{
				http.Redirect(w, r, "/login/password", 301)
			}
		default:
			{

				fmt.Println(" We got this Account ", accounts)
				files = []string{
					app.Path + "/login/accounts.page.html",
				}
				ts, err := template.ParseFiles(files...)
				if err != nil {
					app.ErrorLog.Println(err.Error())
					return
				}
				err = ts.Execute(w, accounts)
				if err != nil {
					app.ErrorLog.Println(err.Error())
				}
			}
		}

	}
}
