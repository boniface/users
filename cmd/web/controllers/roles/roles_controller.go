package roles

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"users/cmd/config"
)

func Roles(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", RolesHanler(app))
	return r

}

func RolesHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			"./views/html/roles/roles.page.html",
			"./views/html/base/base.page.html",
			"./views/html/base/footer.page.html",
			"./views/html/base/header.page.html",
			"./views/html/base/modal.page.html",
			"./views/html/base/navbar.page.html",
			"./views/html/base/sidebar.page.html",
			"./views/html/base/tabs.page.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.ExecuteTemplate(w, "base", nil)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}
