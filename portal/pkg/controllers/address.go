package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/konstfish/whitelist-portal/portal/pkg/templates"
)

func GetAddress(c *gin.Context) {
	user, exists := c.Get("user")
	// should never happen
	if !exists {
		c.AbortWithStatus(401)
		return
	}

	component := templates.ListEntry(user.(string))
	component.Render(c.Request.Context(), c.Writer)
}
