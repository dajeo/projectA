package main

import (
	"projectA/config"
	"projectA/db"
	"projectA/docs"
	"projectA/middlewares"
	"projectA/server"
)

func main() {
	docs.SwaggerInfo.BasePath = "/"

	config.Init("dev")

	db.InitDB()
	db.InitRedis()

	middlewares.InitJWT()

	server.Init()
}
