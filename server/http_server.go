package server

import (
	"log"
	"net/http"

	"go_web_server/config"
)

func StartHttpServer() {
	http.HandleFunc("/", HandleRoot)

	err := http.ListenAndServe(config.ServerPort, nil)

	if err != nil {
		log.Fatal("Server Error:", err)
	}
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Mythos Ma's Go Web Server!"))
}
