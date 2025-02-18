package auth

import (
	"net/http"

	"go_web_server/pkg/response"
)

func CheckTokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		response.Success(w, true, "")
	} else {
		response.Error(w, http.StatusMethodNotAllowed, "请求方法错误", nil)
	}
}
