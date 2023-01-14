package server

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
	"projectA/controllers"
)

func NewRouter(authMiddleware *jwt.GinJWTMiddleware) *gin.Engine {
	r := gin.Default()

	customersController := new(controllers.CustomersController)
	gatewayController := new(controllers.GatewayController)
	gatewayController.Auth = authMiddleware

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/authorize", authMiddleware.LoginHandler)
		authGroup.GET("/refresh", authMiddleware.RefreshHandler)
	}

	customersGroup := r.Group("/customers")
	{
		customersGroup.Use(authMiddleware.MiddlewareFunc())
		{
			customersGroup.GET("/create_task", customersController.CreateTask)
		}
	}

	r.GET("/gateway", gatewayController.Handle)

	r.GET("/swagger/*any", swagger.WrapHandler(files.Handler))

	return r
}
