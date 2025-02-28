package service

import (
	"strings"

	"go_web_server/db"
)

func CreateMarkdownRecord(title string, slug string, content string) (int64, string, error) {

	tx, err := db.DB.Begin()
	if err != nil {
		return 0, "服务器错误", err
	}
	defer tx.Rollback()

	result, err := tx.Exec("INSERT INTO articles (title, slug, content) VALUES (?, ?, ?)", title, slug, content)

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return 0, "slug已存在", err
		}
		return 0, "服务器错误", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, "服务器错误", err
	}

	if err = tx.Commit(); err != nil {
		return 0, "服务器错误", err
	}

	return id, "新增成功", nil
}

func UpdateMarkdownRecord(id int, title string, slug string, content string) (string, error) {
	tx, err := db.DB.Begin()
	if err != nil {
		return "服务器错误", err
	}
	defer tx.Rollback()

	result, err := tx.Exec("UPDATE articles SET title = ?, slug = ?, content = ? WHERE id = ?", title, slug, content, id)

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return "slug已存在", err
		}
		return "服务器错误", err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return "服务器错误", err
	}

	if rowsAffected == 0 {
		return "记录不存在", nil
	}

	return "更新成功", nil
}
