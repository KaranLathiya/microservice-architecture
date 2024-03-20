package routes

import (
	"net/http"
	"portfolio/Auth/middleware"
	"portfolio/handler/portfolio"
	"github.com/go-chi/chi"
)

func InitializeRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {

		r.Route("/expert", func(r chi.Router) {
			r.Use(middleware.Middleware)
			r.Get("/portfolio", portfolio.PortfolioList)
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
