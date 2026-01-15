package handlers

import (
	"fmt"
	"io"
	"microservices/images/files"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

// Files is a handler for reading and writing files
type Files struct {
	log   hclog.Logger
	store files.Storage
}

// NewFiles creates a new File handler
func NewFiles(s files.Storage, l hclog.Logger) *Files {
	return &Files{store: s, log: l}
}

// UploadRest implements the http.Handler interface
func (f *Files) UploadRest(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]

	f.log.Info("Handle POST", "id", id, "filename", fn)

	// check that the file path is a valid name and file
	if id == "" || fn == "" {
		f.invalidURI(r.URL.String(), rw)
	}

	f.saveFile(id, fn, rw, r.Body)
}

// UploadMultipart handles multipart requests
func (f *Files) UploadMultipart(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(128 * 1024)
	if err != nil {
		f.log.Error("Bad Request", "error", err)
		http.Error(rw, "Expected multipart form data", http.StatusBadRequest)

		return
	}

	id := r.FormValue("id")
	_, idErr := strconv.Atoi(id)
	if idErr != nil {
		f.log.Error("Bad request", "error", idErr)
		http.Error(rw, "Expected integer id", http.StatusBadRequest)

		return
	}

	f.log.Info("Process form for", "id", id)

	ff, mh, err := r.FormFile("file")
	fmt.Println(ff, mh, err)
	if err != nil {
		f.log.Error("Bad request", "error", err)
		http.Error(rw, "Expected file", http.StatusBadRequest)

		return
	}

	f.saveFile(id, mh.Filename, rw, ff)
}

func (f *Files) invalidURI(uri string, rw http.ResponseWriter) {
	f.log.Error("Invalid path", "path", uri)
	http.Error(rw, "Invalid file path should be in the format: /[id]/[filepath]", http.StatusBadRequest)
}

// saveFile saves the contents of the request to a file
func (f *Files) saveFile(id string, path string, rw http.ResponseWriter, r io.ReadCloser) {
	f.log.Info("Save file for product", "id", id, "path", path)

	fp := filepath.Join(id, path)
	err := f.store.Save(fp, r)
	if err != nil {
		f.log.Error("Unable to save file", "error", err)
		http.Error(rw, "Unable to save file", http.StatusInternalServerError)
	}
}
