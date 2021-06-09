package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hey I am Alive",
	})
}
