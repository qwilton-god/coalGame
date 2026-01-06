package myhttp

import "github.com/gin-gonic/gin"

func SendError(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{
		"error": err.Error(),
	})
}
