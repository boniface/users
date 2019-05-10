package mail

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"users/cmd/config"
)

func Mail(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", MailHanler(app))
	return r

}

func MailHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			"./views/html/mail/mail.page.html",
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
