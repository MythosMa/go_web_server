package service

import (
	"database/sql"
	"errors"
	"strings"

	"go_web_server/db"
	"go_web_server/pkg/model"
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

func UpdateMarkdownRecord(id int64, title string, slug string, content string) (int64, string, error) {
	tx, err := db.DB.Begin()
	if err != nil {
		return id, "服务器错误", err
	}
	defer tx.Rollback()

	result, err := tx.Exec("UPDATE articles SET title = ?, slug = ?, content = ? WHERE id = ?", title, slug, content, id)

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return id, "slug已存在", err
		}
		return id, "服务器错误", err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return id, "服务器错误", err
	}

	if rowsAffected == 0 {
		return id, "记录不存在", nil
	}

	if err = tx.Commit(); err != nil {
		return 0, "服务器错误", err
	}

	return id, "更新成功", nil
}

func GetMarkdownRecord(id int64) (*model.MarkdownFile, string, error) {
	var markdown model.MarkdownFile

	tx, err := db.DB.Begin()
	if err != nil {
		return nil, "服务器错误", err
	}
	defer tx.Rollback()

	err = tx.QueryRow("SELECT id, title, slug, content, created_at, updated_at FROM articles WHERE id = ?", id).Scan(&markdown.ID, &markdown.Title, &markdown.Slug, &markdown.Content, &markdown.CreatedAt, &markdown.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, "数据不存在", err
		}
		return nil, "服务器错误", err
	}

	return &markdown, "查询成功", nil
}
