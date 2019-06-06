package security

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"users/cmd/config"
	"users/pkg/api/security"
)

func Security(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", SecurityHanler(app))
	return r

}

func SecurityHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		apiKeys, err := security.GetApiKeys()

		if err != nil {
			app.ServerError(w, err)
		}

		siteAccessKeys, err := security.GetSiteAccessKeysApis()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			apiKey             []security.ApiKeys
			siteAccessKeysApis []security.SiteAccessKeysApi
			Name               string
		}
		data := PageData{apiKeys, siteAccessKeys, " "}

		files := []string{
			app.Path + "/security/security.page.html",
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
