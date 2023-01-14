package main

import (
	"github.com/joho/godotenv"
	"projectA/db"
	"projectA/docs"
	"projectA/middlewares"
	"projectA/server"
)

func main() {
	docs.SwaggerInfo.BasePath = "/"

	envErr := godotenv.Load()
	if envErr != nil {
		return
	}

	db.InitDB()
	db.InitRedis()

	auth := middlewares.InitJWTMiddleware()

	server.Init(auth)
}
