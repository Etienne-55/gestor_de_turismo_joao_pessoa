package main

import (
	"projeto_turismo_jp/db"

	"github.com/gin-gonic/gin"
)


func main() {
	db.InitDB()
	server := gin.Default()

	server.Run(":8080")
}

