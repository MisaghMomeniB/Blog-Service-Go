package handler

import (
	"net/http"
    "io/ioutil"
    "path/filepath"
)

// IndexHandler serves the index.html file
func IndexHandler(w http.ResponseWriter, r *http.Request) {
    // Serve the index.html file
    filePath := filepath.Join("static", "index.html")
    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        http.Error(w, "Could not load HTML file", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "text/html")
    w.Write(data)
}
