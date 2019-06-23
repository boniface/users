package sites

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"users/cmd/config"
	sites2 "users/pkg/api/sites"
)

func Sites(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", sitesHandler(app))
	r.Post("/create", createSiteHandler(app))
	r.Post("/update", updateSiteHandler(app))
	return r

}

func sitesHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		sites, err := sites2.GetSites()

		if err != nil {
			app.ServerError(w, err)
		}
		type PageData struct {
			Sites []sites2.Site
		}
		data := PageData{sites}

		files := []string{
			app.Path + "/sites/sites.page.html",
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

func createSiteHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		siteId := r.PostFormValue("siteId")
		siteUrl := r.PostFormValue("siteUrl")
		siteName := r.PostFormValue("siteName")
		siteAdminEmail := r.PostFormValue("siteAdminEmail")
		site := sites2.Site{siteId, siteName, siteUrl, siteAdminEmail}
		_, err := sites2.CreateSite(site)
		if err != nil {
			app.ServerError(w, err)
		}

		fmt.Println(" The Object In Is ", site)
		http.Redirect(w, r, "/sites", 301)
	}
}

func updateSiteHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		r.ParseForm()
		siteId := r.PostFormValue("siteId")
		siteUrl := r.PostFormValue("siteUrl")
		siteName := r.PostFormValue("siteName")
		siteAdminEmail := r.PostFormValue("siteAdminEmail")
		site := sites2.Site{siteId, siteName, siteUrl, siteAdminEmail}
		_, err := sites2.UpdateSite(site)
		if err != nil {
			app.ServerError(w, err)
		}

		http.Redirect(w, r, "/sites", 301)
	}
}
