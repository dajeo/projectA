package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projectA/models"
	"projectA/utils"
)

func Root(context *gin.Context) {
	var tag models.Tag
	utils.Db.First(&tag, context.Query("id"))
	context.JSON(http.StatusOK, tag)
}
