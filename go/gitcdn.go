package main

import (
	"io"
	"mime"
	"net/http"
	"path/filepath"
)

func main() {
	http.HandleFunc(`/`, serveFile)
	if err := http.ListenAndServe(`:80`, nil); err != nil {
		panic(err)
	}
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	fileMime := mime.TypeByExtension(filepath.Ext(r.URL.Path))

	w.Header().Set(`Content-Type`, fileMime)

	fileData, err := http.Get(`https://raw.githubusercontent.com` + r.URL.Path)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	defer fileData.Body.Close()

	_, err = io.Copy(w, fileData.Body)

	if err != nil {
		w.WriteHeader(500)
		return
	}
}
