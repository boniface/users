package subscriptions

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"sort"
	"users/cmd/config"
	"users/pkg/api/sites"
	"users/pkg/api/subscription"
)

func SiteSubscriptions(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", siteSubscriptionsHandler(app))
	return r

}
func siteSubscriptionsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		siteSubs, err := subscription.GetSiteSubscriptions()

		if err != nil {
			app.ServerError(w, err)
		}

		sort.SliceStable(siteSubs, func(i, j int) bool {
			return siteSubs[i].EndDate.After(siteSubs[j].EndDate)
		})

		site, err := sites.GetSites()
		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			SiteSubs []subscription.SiteSubscription
			Sites    []sites.Site
		}
		data := PageData{siteSubs, site}

		files := []string{
			app.Path + "/subscriptions/sitebuscription/site.page.html",
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
