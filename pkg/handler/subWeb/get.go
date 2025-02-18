package subweb

import (
	"net/http"

	"go_web_server/pkg/response"
	"go_web_server/pkg/service"
)

func GetSubWebHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		subWebs, msg, err := service.GetSubWeb()

		if err != nil {
			response.Error(w, http.StatusInternalServerError, msg, err)
			return
		}

		response.Success(w, subWebs, msg)
	} else {
		response.Error(w, http.StatusMethodNotAllowed, "请求方法错误", nil)
	}
}
