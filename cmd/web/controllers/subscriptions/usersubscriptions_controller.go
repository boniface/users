package subscriptions

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"users/cmd/config"
	"users/pkg/api/sites"
	"users/pkg/api/subscription"
)

func UserSubscriptions(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", userSubscriptionsHanler(app))
	return r

}

func userSubscriptionsHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userSubs, err := subscription.GetUserSubscriptions()

		if err != nil {

			app.ServerError(w, err)
		}

		site, err := sites.GetSites()
		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			UserSubs []subscription.UserSubscriptions
			Sites    []sites.Site
		}
		data := PageData{userSubs, site}

		files := []string{
			app.Path + "/subscriptions/usersubscription/usersubs.page.html",
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
