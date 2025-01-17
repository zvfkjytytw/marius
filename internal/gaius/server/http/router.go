package gaiushttp

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *ServerHTTP) newRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.StripSlashes)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	// ping handler
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	r.Post("add_mus", s.addMus)

	r.Route("/save_file", func(r chi.Router) {
		r.Use(s.saveCtx)
		r.Post("/*", s.saveFile)
	})

	r.Route("/get_file", func(r chi.Router) {
		r.Use(s.getCtx)
		r.Get("/*", s.getFile)
	})

	r.Route("/update_file", func(r chi.Router) {
		r.Use(s.getCtx)
		r.Put("/*", s.updateFile)
	})

	r.Route("/delete_file", func(r chi.Router) {
		r.Use(s.getCtx)
		r.Delete("/*", s.deleteFile)
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
