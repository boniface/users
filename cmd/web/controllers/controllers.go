package controllers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"users/cmd/config"
	"users/cmd/web/controllers/login"
	"users/cmd/web/controllers/mail"
	"users/cmd/web/controllers/roles"
	"users/cmd/web/controllers/security"
	"users/cmd/web/controllers/sites"
	"users/cmd/web/controllers/subscriptions"
	"users/cmd/web/controllers/systemlogs"
	"users/cmd/web/controllers/users"
)

func Controllers(env *config.Env) http.Handler {
	mux := chi.NewMux()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)

	mux.Handle("/", login.Home(env))
	mux.Mount("/login", login.Login(env))
	mux.Mount("/mail", mail.Mail(env))
	mux.Mount("/roles", roles.Roles(env))
	mux.Mount("/security", security.Security(env))
	mux.Mount("/sites", sites.Sites(env))
	mux.Mount("/subscriptions", subscriptions.Subscriptions(env))
	mux.Mount("/systemlogs", systemlogs.SitesHanler(env))
	mux.Mount("/users", users.UsersHanler(env))

	fileServer := http.FileServer(http.Dir("./views/assets"))
	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/assets/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Mount("/assets/", http.StripPrefix("/assets", fileServer))
	return mux

}
