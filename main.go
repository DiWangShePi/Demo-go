package main

import (
	"Demo/global"
	"Demo/initialize"
	"fmt"
	"net/http"
	"time"
)

// @title gin+gorm crud 测试swagger
// @version 1.0
// @host localhost:9090
// @BashPath /
func main() {
	initialize.InitDb()
	db, _ := global.GLO_DB.DB()
	defer db.Close()

	Router := initialize.InitRouter()
	server := &http.Server{
		Addr:           "localhost:9090",
		Handler:        Router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
