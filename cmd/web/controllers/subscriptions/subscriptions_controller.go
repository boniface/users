package subscriptions

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"users/cmd/config"
	"users/pkg/api/subscription"
)

func Subscriptions(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", subscriptionsHanler(app))
	return r

}

func subscriptionsHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		subs, err := subscription.GetSubscriptions()
		if err != nil {
			app.ServerError(w, err)
		}

		usesubs, err := subscription.GetUserSubscriptions()

		if err != nil {

			app.ServerError(w, err)
		}

		siteSubs, err := subscription.GetSiteSubscriptions()

		if err != nil {
			app.ServerError(w, err)
		}

		subTypes, err := subscription.GetSubscriptionTypes()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			Subs     []subscription.Subscription
			SubTypes []subscription.SubscriptionTypes
			SiteSubs []subscription.SiteSubscription
			UserSubs []subscription.UserSubscriptions
		}
		data := PageData{subs, subTypes, siteSubs, usesubs}

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
