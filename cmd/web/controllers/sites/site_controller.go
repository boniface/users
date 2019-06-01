package sites

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"users/cmd/config"
	"users/pkg/api/roles"
	sites2 "users/pkg/api/sites"
)

func Sites(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", sitesHanler(app))
	r.Get("/details/{id}", sitesHanler(app))
	r.Get("/edit/{id}", sitesHanler(app))
	r.Post("/create", sitesHanler(app))
	r.Post("/update", sitesHanler(app))
	return r

}

func sitesHanler(app *config.Env) http.HandlerFunc {
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

func siteEditHanler(app *config.Env) http.HandlerFunc {
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

func sitesDetailsHanler(app *config.Env) http.HandlerFunc {
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

func createHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		Id := r.PostFormValue("id")
		RoleName := r.PostFormValue("roleName")
		Description := r.PostFormValue("description")
		role := roles.RolesPool{Id, RoleName, Description}
		_, err := roles.CreateRolespool(role)
		if err != nil {
			app.ServerError(w, err)
		}
		http.Redirect(w, r, "/roles", 301)
	}
}

func updateHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		Id := r.PostFormValue("id")
		RoleName := r.PostFormValue("roleName")
		Description := r.PostFormValue("description")
		role := roles.RolesPool{Id, RoleName, Description}
		_, err := roles.CreateRolespool(role)
		if err != nil {
			app.ServerError(w, err)
		}
		http.Redirect(w, r, "/roles", 301)
	}
}
