package service

import (
	"go_web_server/db"
	"go_web_server/pkg/model"
)

func GetSubWeb() ([]model.SubWeb, string, error) {
	var subWebs []model.SubWeb

	tx, err := db.DB.Begin()
	if err != nil {
		return nil, "服务器错误", err
	}
	defer tx.Rollback()

	rows, err := tx.Query("SELECT * FROM sub_web")

	if err != nil {
		return nil, "服务器错误", err
	}

	// 检查是否有错误
	if err := rows.Err(); err != nil {
		return nil, "数据查询错误", err
	}

	// 遍历结果集
	for rows.Next() {
		var subWeb model.SubWeb
		err := rows.Scan(&subWeb.ID, &subWeb.Name, &subWeb.Url)
		if err != nil {
			return nil, "数据查询错误", err
		}
		subWebs = append(subWebs, subWeb)
	}

	if err != nil {
		return nil, "服务器错误", err
	}

	return subWebs, "获取成功", nil
}
