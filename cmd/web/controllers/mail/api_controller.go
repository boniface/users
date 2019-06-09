package mail

import (
	"github.com/go-chi/chi"
	"net/http"
	"users/cmd/config"
)

func MailApi(app *config.Env) http.Handler {
	r := chi.NewRouter()

	return r
}
