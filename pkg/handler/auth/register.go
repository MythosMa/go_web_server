package auth

import (
	"encoding/json"
	"net/http"

	"go_web_server/pkg/model"
	"go_web_server/pkg/response"
	"go_web_server/pkg/service"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var user model.User

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			response.Error(w, http.StatusBadRequest, "请求参数错误", err)
			return
		}

		createdUser, msg, err := service.RegisterUser(user.Username, user.Password, user.Email)

		if err != nil {
			response.Error(w, http.StatusInternalServerError, msg, err)
			return
		}

		response.Success(w, createdUser.ToUserResponse(), msg)
	} else {
		response.Error(w, http.StatusMethodNotAllowed, "请求方法错误", nil)
	}
}
