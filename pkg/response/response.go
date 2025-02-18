package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JSON(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Set(("content-type"), "application/json")

	resp := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}

	jsonBytes, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{
			Code:    http.StatusInternalServerError,
			Message: "服务器错误",
			Data:    nil,
		})
		return
	}

	w.WriteHeader(code)
	w.Write(jsonBytes)
}

func Success(w http.ResponseWriter, data interface{}, message ...string) {
	msg := "success"
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	JSON(w, http.StatusOK, msg, data)
}

func Error(w http.ResponseWriter, code int, message string, err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	JSON(w, code, message, nil)
}
