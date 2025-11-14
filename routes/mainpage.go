package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func mainPage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "This BlogString"})
}
