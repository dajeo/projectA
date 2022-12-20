package auth

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	g := r.Group("auth")
	{
		g.GET("/", Root)
	}
}
