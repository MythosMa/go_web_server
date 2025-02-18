package test

import (
	"net/http"

	"go_web_server/pkg/response"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		response.Success(w, nil, "测试token成功")
	} else {
		response.Error(w, http.StatusMethodNotAllowed, "请求方法错误", nil)
	}
}
