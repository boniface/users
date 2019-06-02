package systemlogs

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"users/cmd/config"
	sites2 "users/pkg/api/sites"
	"users/pkg/api/systemlogs"
)

func Systemlogs(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", logsHandler(app))
	r.Get("/details/{id}", logsForSiteHandler(app))
	r.Get("/site/{siteid}", logsDetailsHandler(app))
	r.Post("/logsites", logsSiteHandler(app))
	return r

}

func logsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		logs, err := systemlogs.GetLogEvents()

		if err != nil {
			app.ServerError(w, err)
		}

		sites, err := sites2.GetSites()

		if err != nil {
			app.ServerError(w, err)
		}
		type PageData struct {
			Sites []sites2.Site
			Logs  []systemlogs.LogEvent
		}
		data := PageData{sites, logs}

		files := []string{
			app.Path + "/systemlogs/systemlogs.page.html",
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

func logsForSiteHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		siteId := chi.URLParam(r, "siteid")
		logs, err := systemlogs.GetLogEventsForSite(siteId)

		if err != nil {
			app.ServerError(w, err)
		}

		sites, err := sites2.GetSites()

		if err != nil {
			app.ServerError(w, err)
		}
		type PageData struct {
			Sites []sites2.Site
			Logs  []systemlogs.LogEvent
		}
		data := PageData{sites, logs}

		files := []string{
			app.Path + "/systemlogs/systemlogs.page.html",
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

func logsDetailsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := chi.URLParam(r, "id")
		log, err := systemlogs.GetLogEvent(id)

		if err != nil {
			app.ServerError(w, err)
		}

		sites, err := sites2.GetSites()

		if err != nil {
			app.ServerError(w, err)
		}
		type PageData struct {
			Sites []sites2.Site
			Log   systemlogs.LogEvent
		}
		data := PageData{sites, log}

		files := []string{
			app.Path + "/systemlogs/logevents/details.page.html",
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

func logsSiteHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		siteId := r.PostFormValue("siteId")
		//RoleName := r.PostFormValue("roleName")
		//Description := r.PostFormValue("description")
		//role := roles.RolesPool{Id, RoleName, Description}
		//_, err := roles.UpdateRolespool(role)
		//if err != nil {
		//	app.ServerError(w, err)
		//}

		fmt.Println(" The Site ID is ", siteId)
		http.Redirect(w, r, "/logs", 301)

	}
}
