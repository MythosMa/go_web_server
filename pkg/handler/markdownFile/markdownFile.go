package markdownfile

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go_web_server/pkg/response"
	"go_web_server/pkg/service"
)

type CreateMarkdownFileRequestInfo struct {
	Title   string `json:"title"`
	Slug    string `json:"slug"`
	Content string `json:"content"`
}

type CreateMarkdownFileResponseInfo struct {
	ID int64 `json:"id"`
}

type UpdateMarkdownFileRequestInfo struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Slug    string `json:"slug"`
	Content string `json:"content"`
}

type UpdateMarkdownFileResponseInfo struct {
	ID int64 `json:"id"`
}

type GetMarkdownFileResponseInfo struct {
	ID int64 `json:"id"`
}

func CreateMarkdownFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var info CreateMarkdownFileRequestInfo

		if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
			response.Error(w, http.StatusBadRequest, "请求参数错误", err)
			return
		}

		id, msg, err := service.CreateMarkdownRecord(info.Title, info.Slug, info.Content)

		if err != nil {
			response.Error(w, http.StatusInternalServerError, msg, err)
			return
		}

		var responseInfo CreateMarkdownFileResponseInfo
		responseInfo.ID = id

		response.Success(w, responseInfo, msg)

	} else {
		response.Error(w, http.StatusMethodNotAllowed, "请求方法错误", nil)
	}
}

func UpdateMarkdownFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var info UpdateMarkdownFileRequestInfo

		if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
			response.Error(w, http.StatusBadRequest, "请求参数错误", err)
			return
		}

		id, msg, err := service.UpdateMarkdownRecord(info.ID, info.Title, info.Slug, info.Content)

		if err != nil {
			response.Error(w, http.StatusInternalServerError, msg, err)
			return
		}

		var responseInfo UpdateMarkdownFileResponseInfo
		responseInfo.ID = id

		response.Success(w, responseInfo, msg)
	} else {
		response.Error(w, http.StatusMethodNotAllowed, "请求方法错误", nil)
	}
}

func GetMarkdownFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idStr := r.URL.Query().Get("id")

		if idStr == "" {
			response.Error(w, http.StatusBadRequest, "请求参数错误", nil)
			return
		}

		id, err := strconv.ParseInt(idStr, 10, 64)

		if err != nil {
			response.Error(w, http.StatusBadRequest, "请求参数错误", err)
			return
		}

		markdown, msg, err := service.GetMarkdownRecord(id)

		if err != nil {
			response.Error(w, http.StatusInternalServerError, msg, err)
			return
		}

		if markdown == nil {
			response.Error(w, http.StatusNotFound, msg, nil)
			return
		}

		response.Success(w, markdown, msg)

	} else {
		response.Error(w, http.StatusMethodNotAllowed, "请求方法错误", nil)
	}

}
