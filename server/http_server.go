package server

import (
	"log"
	"net/http"
	"os"
)

func StartHttpServer() {

	http.HandleFunc("/", HandleRoot)

	port := os.Getenv("PORT")
	err := http.ListenAndServe("0.0.0.0:"+port, nil)

	if err != nil {
		log.Fatal("Server Error:", err)
	}
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Mythos Ma's Go Web Server!"))
}
