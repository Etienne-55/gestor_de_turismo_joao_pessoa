package main

import (
	"projeto_turismo_jp/db"
	"projeto_turismo_jp/routes"
	"projeto_turismo_jp/server"

)


func main() {
	db.InitDB()
	server := server.SetupServer()

	routes.AppRoutes(server)

	server.Run(":8080")
}

