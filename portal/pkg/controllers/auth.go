package controllers

import (
	"github.com/gin-gonic/gin"
)

func ValidateAuthHeader(c *gin.Context) {
	header := c.GetHeader(AuthHeader)

	if header == "" {
		c.AbortWithStatus(401)
		return
	}

	c.Set("user", header)

	c.Next()
}
