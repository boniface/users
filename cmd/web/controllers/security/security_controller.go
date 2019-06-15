package security

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"users/cmd/config"
	"users/pkg/api/security"
)

func Security(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", securityHandler(app))
	return r

}

func securityHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		siteAccessKeys, err := security.GetSiteAccessKeysApis()

		if err != nil {
			app.ServerError(w, err)
		}

		apiKeys, err := security.GetApiKeys()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			ApiKey  []security.ApiKeys
			SiteKey []security.SiteAccessKeysApi
			Name    string
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
			fmt.Println(" it Breaks here 3")
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.ExecuteTemplate(w, "base", data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}
