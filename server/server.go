package server

import (
	"log"
	"net/http"

	"ascii-art-web/handlers"
)

func RegisterHandlers() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/export", ExportHandler) // Use the correct package and function name
	http.HandleFunc("/", handlers.ServeTemplate)
	http.HandleFunc("/ascii-art", handlers.HandleAsciiArt)
}

func StartServer() {
	RegisterHandlers()

	log.Print("Listening on :http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
