package customers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	g := r.Group("customers")

	g.Use(authMiddleware.MiddlewareFunc())
	{
		g.GET("create_task", CreateTask)
	}
}
