package chi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func getArticle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func UserRouter() *chi.Mux {
	apiRouter := chi.NewRouter()
	apiRouter.Get("/articles/{date}-{slug}", getArticle)

	return apiRouter
}
