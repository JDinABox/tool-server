package toolserver

import (
	"bytes"
	"log"
	"net/http"
	"strings"

	"github.com/JDinABox/tool-server/internal/adguard"
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

	// file server
	r.Handle("/f/*", http.StripPrefix("/f", http.FileServer(http.Dir("/var/lib/tool-server/files"))))

	// lists
	adguardServer := adguard.New()

	// /l
	r.Route("/l", func(r chi.Router) {
		// /l/services
		r.Route("/services", func(r chi.Router) {
			// /l/services
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				services, err := adguardServer.AdguardContent(r.Context())
				if err != nil {
					log.Println(err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				serviceList := services.MappedNames()

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(map[string]any{
					"services": serviceList,
				})
			})

			// /l/services/{service(s)}
			// service(s) can be a + separated list of services
			r.Get("/{services}", func(w http.ResponseWriter, r *http.Request) {
				serviceStr := chi.URLParam(r, "services")
				serviceStr = strings.ToLower(strings.TrimSpace(serviceStr))
				services := strings.Split(serviceStr, "+")

				knownServices, err := adguardServer.AdguardContent(r.Context())
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				var output bytes.Buffer
				for _, v := range services {
					service := strings.TrimSpace(v)
					if service == "" {
						continue
					}
					s, err := knownServices.ServiceList(service)
					if err != nil {
						w.WriteHeader(http.StatusNotFound)
						json.NewEncoder(w).Encode(
							map[string]any{
								"error":     "404",
								"not-found": service,
							},
						)
						return
					}
					output.WriteString(s)
				}
				w.Write(output.Bytes())
			})
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
