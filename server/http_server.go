package server

import (
	"log"
	"net/http"
	"os"

	"go_web_server/pkg/handler/auth"
	"go_web_server/pkg/handler/test"
	"go_web_server/pkg/middleware"
)

func StartHttpServer() {

	http.HandleFunc("/", HandleRoot)
	http.HandleFunc("/register", auth.RegisterHandler)
	http.HandleFunc("/login", auth.LoginHandler)
	http.Handle("/test", middleware.AuthMiddleware(http.HandlerFunc(test.TestHandler)))

	port := os.Getenv("PORT")
	err := http.ListenAndServe("0.0.0.0:"+port, nil)

	if err != nil {
		log.Fatal("Server Error:", err)
	}
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Mythos Ma's Go Web Server!"))
}
