package config

import (
	"la-blog-go/db"
	"la-blog-go/router"
)

func Init() {
	// 初始化数据库
	db.ConnectToSQLiteDB("SQLite.db", true)
	// 初始化路由
	router.InitRouters()
}
