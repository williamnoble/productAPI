package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"path/filepath"
	"productAPI/fs"
)

type FileHandler struct{
	log *log.Logger
	store fs.FileStore
}

func NewFileHandler(l *log.Logger, s fs.FileStore) *FileHandler {
	return &FileHandler{
		log:   l,
		store: s,
	}
}

func (f *FileHandler) Upload(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	filename := vars["filename"]
	reader := r.Body

	path := filepath.Join(id, filename)
	err := f.store.Save(path, reader)
	if err != nil {
		f.log.Println("Error", err)
		http.Error(w, "error uploading", http.StatusInternalServerError)
	}
}


