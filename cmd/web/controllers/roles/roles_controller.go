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
	//Roles
	r.Post("/create", CreateRoleHandler(app))
	r.Post("/update", UpdateRoleHandler(app))
	r.Get("/edit/{id}", EditRoleHandler(app))
	r.Get("/details/{id}", DetailsRoleHandler(app))
	//Roles Pool
	r.Post("/pool/create", CreatePoolRoleHandler(app))
	r.Post("/pool/update", UpdatePoolRoleHandler(app))
	r.Get("/pool/edit/{id}", EditPoolRoleHandler(app))
	r.Get("/pool/details/{id}", DetailPoolRoleHandler(app))
	return r

}

//SiteId      string `json:"siteId"`
//Id          string `json:"id"`
//RoleName    string `json:"roleName"`
//Description

func CreatePoolRoleHandler(app *config.Env) http.HandlerFunc {
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

func EditPoolRoleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		//email := r.PostFormValue("email")
		//chi.URLParam(r, "articleID");

	}

}

func DetailPoolRoleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		//email := r.PostFormValue("email")
		//chi.URLParam(r, "articleID");

	}

}

func EditRoleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		//email := r.PostFormValue("email")
		//chi.URLParam(r, "articleID");

	}

}

func DetailsRoleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		//email := r.PostFormValue("email")
		//chi.URLParam(r, "articleID");

	}

}

func UpdatePoolRoleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//r.ParseForm()
		//email := r.PostFormValue("email")

	}

}

func CreateRoleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		//email := r.PostFormValue("email")

	}

}

func UpdateRoleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		//email := r.PostFormValue("email")

	}

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
