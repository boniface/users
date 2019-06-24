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
	//Roles Details
	r.Post("/details/{id}", DetailsRoleHandler(app))
	//Roles Pool
	r.Post("/pool/create", CreatePoolRoleHandler(app))
	r.Post("/pool/update", UpdatePoolRoleHandler(app))
	r.Get("/edit/{id}", EditPoolRoleHandler(app))
	r.Get("/details/{id}", DetailPoolRoleHandler(app))
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
		id := chi.URLParam(r, "id")
		role, _ := roles.GetRolespool(id)
		type PageData struct {
			Rolepool roles.RolesPool
		}
		data := PageData{role}
		files := []string{
			app.Path + "/roles/rolespool/usersubs.page.html",
			app.Path + "/base/header.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebar.page.html",
			app.Path + "/base/footer.page.html",
			app.Path + "/base/modal.page.html",
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

func DetailPoolRoleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		role, _ := roles.GetRolespool(id)
		type PageData struct {
			Rolepool roles.RolesPool
		}
		data := PageData{role}
		files := []string{
			app.Path + "/roles/rolespool/site.page.html",
			app.Path + "/base/header.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebar.page.html",
			app.Path + "/base/footer.page.html",
			app.Path + "/base/modal.page.html",
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

func DetailsRoleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		role, _ := roles.GetRolespool(id)
		type PageData struct {
			Rolepool roles.RolesPool
		}
		data := PageData{role}
		files := []string{
			app.Path + "/roles/roles.page.html",
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

func UpdatePoolRoleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		Id := r.PostFormValue("id")
		RoleName := r.PostFormValue("roleName")
		Description := r.PostFormValue("description")
		role := roles.RolesPool{Id, RoleName, Description}
		_, err := roles.UpdateRolespool(role)
		if err != nil {
			app.ServerError(w, err)
		}
		http.Redirect(w, r, "/roles", 301)
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
			app.Path + "/roles/roles.page.html",
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
