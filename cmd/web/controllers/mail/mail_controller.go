package mail

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"users/cmd/config"
	"users/pkg/api/mail"
)

func Mail(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", MailHanler(app))
	return r

}

func MailHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		mailapi, err := mail.GetMailApis()

		if err != nil {
			app.ServerError(w, err)
		}

		mailconfig, err := mail.GetMailconfigs()

		if err != nil {
			app.ServerError(w, err)
		}

		smtpconfig, err := mail.GetSmtpConfigs()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			Mail    []mail.MailApi
			MailConfig []mail.MailConfig
			SmtpConfig []mail.SmtpConfig
			Name     string
		}

		data := PageData{mailapi,mailconfig, smtpconfig,""}

		files := []string{
			app.Path + "/mail/mail.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebar.page.html",
			app.Path + "/base/footer.page.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.ExecuteTemplate(w, "base", data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}
