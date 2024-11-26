package gaiushttp

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (h *ServerHTTP) newRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.StripSlashes)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	// ping handler
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	// stubs
	r.Delete("/*", notImplementedYet)
	r.Get("/*", notImplementedYet)
	r.Post("/*", notImplementedYet)
	r.Put("/*", notImplementedYet)

	return r
}

// not implemented handlers.
func notImplementedYet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not implemented yet"))
}
