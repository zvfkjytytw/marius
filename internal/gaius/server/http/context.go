package gaiushttp

import (
	"context"
	"net/http"
	"strings"
)

type contextKey int8

const (
	slash                    = "/"
	contextFileID contextKey = iota
	contextFilePath
)

func (s *ServerHTTP) saveCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filePath := strings.Split(strings.TrimLeft(r.URL.Path, slash), slash)[1:]
		if len(filePath) == 0 {
			http.Error(w, "file not specified", http.StatusBadRequest)
			return
		}

		_, err := s.fs.FindFile(strings.Join(filePath, slash))
		if err == nil {
			http.Error(w, "file already exists", http.StatusBadRequest)
			return
		}

		fileID, err := s.fs.CreateFile(strings.Join(filePath, slash))
		if err != nil {
			s.logger.Sugar().Errorf("failed create file: %v", err)
			http.Error(w, "failed create file", http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), contextFileID, fileID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *ServerHTTP) getCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filePath := strings.Split(r.URL.Path, slash)[1:]
		if len(filePath) == 0 {
			http.Error(w, "file not specified", http.StatusBadRequest)
			return
		}

		fileID, err := s.fs.FindFile(strings.Join(filePath, slash))
		if err != nil {
			http.Error(w, "file not found", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), contextFileID, fileID)
		ctx = context.WithValue(ctx, contextFilePath, strings.Join(filePath, slash))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
