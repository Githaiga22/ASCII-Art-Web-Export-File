package server

import (
	"net/http"
	"strconv"
	"time"
)
func ExportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	asciiArt := r.URL.Query().Get("ascii")
	if asciiArt == "" {
		http.Error(w, "No Art to export", http.StatusBadRequest)
		return
	}

	filename := "ascii_art_" + time.Now().Format("20060102150405") + ".txt"
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Length", strconv.Itoa(len(asciiArt)))
	w.Write([]byte(asciiArt)) // Write the ASCII art to the file
}
