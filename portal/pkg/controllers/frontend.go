package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/konstfish/whitelist-portal/portal/pkg/templates"
)

func GetIndex(c *gin.Context) {
	out, err := templates.HomePage().Render()
	if err != nil {
		c.String(500, "oops")
		c.Header("Refresh", "2")
		return
	}

	c.Writer.Write([]byte(out))
}
