package subscriptions

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"users/cmd/config"

	subscriptions2 "users/pkg/api/subscription"
)

func Subscriptions(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", SubscriptionsHanler(app))
	return r

}

func SubscriptionsHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		subscriptions, err:= subscriptions2.GetSubscriptions()

		if err != nil {
			app.ServerError(w, err)
		}
		type PageData struct {
			Subscription []subscriptions2.Subscription
		}
		data := PageData{subscriptions}

		files := []string{
			app.Path + "/subscriptions/subscriptions.page.html",
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
