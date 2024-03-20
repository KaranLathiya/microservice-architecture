package routes

import (
	"expert/handler/expert"
	"net/http"

	"github.com/go-chi/chi"
)

func InitializeRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {

		r.Route("/user", func(r chi.Router) {
			r.Get("/expert", expert.ExpertList)
		})

		r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(405)
			w.Write([]byte("wrong method"))
		})
		r.NotFound(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(404)
			w.Write([]byte("route does not exist"))
		})

	})
	return r
}
