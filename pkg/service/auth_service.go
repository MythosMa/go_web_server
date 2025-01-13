package service

import (
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"go_web_server/db"
)

func RegisterUser(username, password, email string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "服务器错误", err
	}

	_, err = db.DB.Exec("INSERT INTO users (username, password, avatar) VALUES (?, ?, ?)", username, string(hashedPassword), email)

	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), "Duplicate entry") {
			return "用户已存在", err
		} else {
			return "服务器错误", err
		}
	}

	return "用户注册成功", nil
}
