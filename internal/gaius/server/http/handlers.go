package gaiushttp

import (
	"fmt"
	"net/http"
)

func (s *ServerHTTP) saveFile(w http.ResponseWriter, r *http.Request) {
	fileID := fmt.Sprintf("%v", r.Context().Value(contextFileID))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("file ID: %v", fileID)))
}

func (s *ServerHTTP) getFile(w http.ResponseWriter, r *http.Request) {
	fileID := fmt.Sprintf("%v", r.Context().Value(contextFileID))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("file ID: %v", fileID)))
}

func (s *ServerHTTP) updateFile(w http.ResponseWriter, r *http.Request) {
	fileID := fmt.Sprintf("%v", r.Context().Value(contextFileID))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("file ID: %v", fileID)))
}

func (s *ServerHTTP) deleteFile(w http.ResponseWriter, r *http.Request) {
	fileID := fmt.Sprintf("%v", r.Context().Value(contextFileID))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("file ID: %v", fileID)))
}
