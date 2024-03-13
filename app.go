package toolserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// newApp init fiber app
func newApp() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)

	// Api base path
	r.Route("/api/v1", func(r chi.Router) {
		r.Use(func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				h.ServeHTTP(w, r)
			})
		})

		// ping
		r.Route("/ping", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				json.NewEncoder(w).Encode("pong")
			})
			r.Get("/pop", func(w http.ResponseWriter, r *http.Request) {
				json.NewEncoder(w).Encode("meow")
			})
			r.Get("/meow", func(w http.ResponseWriter, r *http.Request) {
				json.NewEncoder(w).Encode("pop")
			})
		})

		// Return client ip
		r.Get("/ip", func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(map[string]any{"ip": r.RemoteAddr})
		})
	})

	// 404 error
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			map[string]any{
				"error":     "404",
				"not-found": r.URL.String(),
			})
	})

	return r
}
