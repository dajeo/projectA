package customers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateTask godoc
// @Summary Create task
// @Schemes
// @Description do ping
// @Tags Customers
// @Router /customers/create_task [get]
func CreateTask(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(http.StatusOK, gin.H{
		"tel": claims["tel"],
	})
}
