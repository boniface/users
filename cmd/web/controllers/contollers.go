package controllers

import (
	"github.com/go-chi/chi"
	"net/http"
	"users/cmd/config"
)

func Controllers(env *config.Env) http.Handler {
	mux := chi.NewRouter()

	return mux

}
