package service

import (
	"strings"

	"golang.org/x/crypto/bcrypt"

	"go_web_server/db"
	"go_web_server/pkg/model"
)

func RegisterUser(username, password, email string) (*model.User, string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "服务器错误", err
	}

	tx, err := db.DB.Begin()
	if err != nil {
		return nil, "服务器错误", err
	}
	defer tx.Rollback()

	result, err := tx.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?)", username, string(hashedPassword), email)

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			if strings.Contains(err.Error(), "username") {
				return nil, "用户名已存在", err
			} else if strings.Contains(err.Error(), "email") {
				return nil, "邮箱已存在", err
			}
			return nil, "用户注册失败", err
		} else {
			return nil, "服务器错误", err
		}
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, "服务器错误", err
	}

	if err = tx.Commit(); err != nil {
		return nil, "服务器错误", err
	}

	user := &model.User{
		ID:       int(id),
		Username: username,
		Email:    email,
	}

	return user, "用户注册成功", nil
}
