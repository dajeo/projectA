package customers

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	g := r.Group("customers")
	{
		g.GET("create_task", CreateTask)
	}
}
