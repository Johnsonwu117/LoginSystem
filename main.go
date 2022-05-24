package main

import (
	database "LoginSystem/database"
	rou "LoginSystem/router"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	rou.AddUserRouter(v1)
	go func() {
		database.DD()
	}()
	router.Run(":1388")
}

//測試git變更
