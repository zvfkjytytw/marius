package gaiushttp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func (s *ServerHTTP) saveFile(w http.ResponseWriter, r *http.Request) {
	contentType, ok := r.Header["Content-Type"]
	if !ok || contentType[0] != "application/octet-stream" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong Content-Type. Expect application/octet-stream"))
		return
	}

	defer r.Body.Close()
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("failed to read file"))
		return
	}

	fileID, err := strconv.Atoi(fmt.Sprintf("%v", r.Context().Value(contextFileID)))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("wrong new file ID"))
		return
	}

	if err = s.storage.SaveFile(r.Context(), int32(fileID), requestBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to save new file"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("file ID: %v", fileID)))
}

func (s *ServerHTTP) getFile(w http.ResponseWriter, r *http.Request) {
	fileID, err := strconv.Atoi(fmt.Sprintf("%v", r.Context().Value(contextFileID)))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("wrong new file ID"))
		return
	}

	fileData, err := s.storage.GetFile(r.Context(), int32(fileID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to get file data"))
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
	w.Write(fileData)
}

func (s *ServerHTTP) updateFile(w http.ResponseWriter, r *http.Request) {
	contentType, ok := r.Header["Content-Type"]
	if !ok || contentType[0] != "application/octet-stream" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong Content-Type. Expect application/octet-stream"))
		return
	}

	defer r.Body.Close()
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("failed to read file"))
		return
	}

	fileID, err := strconv.Atoi(fmt.Sprintf("%v", r.Context().Value(contextFileID)))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("wrong file ID"))
		return
	}

	if err = s.storage.UpdateFile(r.Context(), int32(fileID), requestBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to update file data"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("file ID: %v", fileID)))
}

func (s *ServerHTTP) deleteFile(w http.ResponseWriter, r *http.Request) {
	fileID, err := strconv.Atoi(fmt.Sprintf("%v", r.Context().Value(contextFileID)))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("wrong new file ID"))
		return
	}
	if err = s.storage.DeleteFile(r.Context(), int32(fileID)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to delete file data"))
		return
	}

	if err = s.fs.DeleteFile(fmt.Sprintf("%v", r.Context().Value(contextFilePath))); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to delete file"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("file has been deleted"))
}

func (s *ServerHTTP) addMus(w http.ResponseWriter, r *http.Request) {
	contentType, ok := r.Header["Content-Type"]
	if !ok || contentType[0] != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong Content-Type. Expect application/json"))
		return
	}

	defer r.Body.Close()
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("failed to read body"))
		return
	}

	mus := struct {
		Address string `json:"address"`
	}{}
	err = json.Unmarshal(requestBody, &mus)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("failed to unmarshal body"))
		return
	}

	if err := s.storage.AddMus(r.Context(), mus.Address); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to add new mus"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("new mus added"))
}
