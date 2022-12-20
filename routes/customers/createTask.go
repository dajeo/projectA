package customers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateTask godoc
// @Summary Create task
// @Schemes
// @Description do ping
// @Tags Customers
// @Router /customers/create_task [get]
func CreateTask(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"a": 1,
	})
}
