package main

import (
	"go-back/model"
	"go-back/router"
	"log"
)

func main() {

	//创建数据库连接
	var conn = model.InitDB()

	defer func() {
		db, _ := conn.DB()
		err := db.Close()
		if err != nil {
			log.Fatal("failed to close connection")
		}
	}()

	router.Start()

}
