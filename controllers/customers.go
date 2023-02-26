package controllers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CustomersController struct{}

// CreateTask godoc
// @Summary Create task
// @Schemes
// @Description do ping
// @Tags Customers
// @Router /customers/create_task [get]
func (controller CustomersController) CreateTask(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(http.StatusOK, gin.H{
		"id": claims["id"],
	})
}
