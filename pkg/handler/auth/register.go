package auth

import (
	"net/http"

	"go_web_server/pkg/service"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("email")

		msg, err := service.RegisterUser(username, password, email)

		if err != nil {
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
