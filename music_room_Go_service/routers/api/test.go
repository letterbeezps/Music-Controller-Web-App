package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestGin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "TestGin",
	})
}
