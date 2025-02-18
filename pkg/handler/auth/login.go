package auth

import (
	"encoding/json"
	"net/http"

	"go_web_server/pkg/jwt"
	"go_web_server/pkg/model"
	"go_web_server/pkg/response"
	"go_web_server/pkg/service"
)

// 登录请求的结构体
type LoginRequestInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponseInfo struct {
	Token string             `json:"token"`    // JWT Token
	User  model.UserResponse `json:"userInfo"` // 用户信息
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var info LoginRequestInfo

		if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
			response.Error(w, http.StatusBadRequest, "请求参数错误", err)
			return
		}

		loginUser, msg, err := service.Login(info.Username, info.Password)

		if err != nil {
			response.Error(w, http.StatusInternalServerError, msg, err)
			return
		}

		token, err := jwt.GenerateToken(loginUser.ID, loginUser.Username)

		if err != nil {
			response.Error(w, http.StatusInternalServerError, "服务器错误", err)
			return
		}

		// 将 Token 和用户信息合并
		loginResponse := LoginResponseInfo{
			Token: token,
			User:  loginUser.ToUserResponse(),
		}

		response.Success(w, loginResponse, msg)
	} else {
		response.Error(w, http.StatusMethodNotAllowed, "请求方法错误", nil)
	}
}
