package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	files "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
	"projectA/docs"
	"projectA/gateway"
	"projectA/routes/auth"
	"projectA/routes/customers"
	"projectA/utils"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		return
	}

	// Initializing databases
	utils.InitDb()
	utils.InitRedis()

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

	// See this before deploy to production
	// https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies
	proxyErr := r.SetTrustedProxies(nil)
	if proxyErr != nil {
		return
	}

	// Initializing routes
	auth.Router(r)
	customers.Router(r)

	// Initialize gateway
	gateway.Router(r)

	r.GET("/swagger/*any", swagger.WrapHandler(files.Handler))
	r.Use(cors.Default()) // Configure before production
	err := r.Run(":8000")
	if err != nil {
		return
	}
}
