package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")

	// 构建 DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error

	// 重试连接数据库
	for i := 0; i < 10; i++ {
		log.Printf("正在尝试连接到MySQL数据库: %v", dsn)
		DB, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Printf("Failed to connect to MySQL (attempt %d): %v", i+1, err)
			time.Sleep(2 * time.Second)
			continue
		}
		err = DB.Ping()
		if err == nil {
			return nil
		}
		log.Printf("Failed to ping MySQL (attempt %d): %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	return nil
}
