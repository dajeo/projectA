package server

import (
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
	"projectA/controllers"
	"projectA/middlewares"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	customersController := new(controllers.CustomersController)
	gatewayController := new(controllers.GatewayController)
	authMiddleware := middlewares.GetJWT()

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
