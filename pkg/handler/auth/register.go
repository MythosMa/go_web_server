package auth

import "net/http"

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("email")
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
