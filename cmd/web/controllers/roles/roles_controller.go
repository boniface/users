package roles

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"users/cmd/config"
	"users/pkg/api/roles"
)

func Roles(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", RolesHanler(app))
	return r

}

func RolesHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		allRoles, err := roles.GetRoles()

		if err != nil {
			app.ServerError(w, err)
		}

		rolespool, err := roles.GetRolespools()

		if err != nil {
			app.ServerError(w, err)
		}
		type PageData struct {
			Roles    []roles.Role
			Rolepool []roles.RolesPool
			Name     string
		}
		data := PageData{allRoles, rolespool, " "}

		files := []string{
			"./views/html/roles/roles.page.html",
			"./views/html/base/base.page.html",
			"./views/html/base/navbar.page.html",
			"./views/html/base/sidebar.page.html",
			"./views/html/base/footer.page.html",
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
