package users

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"users/cmd/config"
	"users/pkg/api/sites"
	"users/pkg/api/users"
)

func Users(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", UsersHanler(app))
	return r

}

func UsersHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		_users, err := users.GetUsers()

		if err != nil {
			app.ServerError(w, err)
		}

		site, err := sites.GetSites()
		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			Users []users.User
			Sites []sites.Site
		}

		data := PageData{_users, site}

		files := []string{
			app.Path + "/users/users.page.html",
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
