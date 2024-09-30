package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/konstfish/whitelist-portal/portal/pkg/cache"
	"github.com/konstfish/whitelist-portal/portal/pkg/templates"
)

func userFromContext(c *gin.Context) string {
	user, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(401)
		return ""
	}
	return user.(string)
}

func GetAddress(c *gin.Context) {
	// user := userFromContext(c)

	component := templates.ListEntry(cache.AddressListEntry{})
	component.Render(c.Request.Context(), c.Writer)
}

func AddAddress(c *gin.Context) {
	user := userFromContext(c)

	var address cache.AddressListEntry
	if err := c.Bind(&address); err != nil {
		c.AbortWithStatus(400)
		return
	}

	if valid, reason := validateAddress(address); !valid {
		component := templates.ErrorMessage(reason)
		component.Render(c.Request.Context(), c.Writer)
		return
	}

	if err := cache.AddAddress(c.Request.Context(), user, address); err != nil {
		c.AbortWithStatus(500)
		return
	}
}

func DeleteAddress(c *gin.Context) {
	user := userFromContext(c)
	address := c.Param("address")
	if address == "" {
		c.AbortWithStatus(400)
		return
	}

	if err := cache.DeleteAddress(c.Request.Context(), user, address); err != nil {
		c.AbortWithStatus(500)
		return
	}
}

func GetAddressList(c *gin.Context) {
	user := userFromContext(c)

	addresses, err := cache.GetAddressList(c.Request.Context(), user)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	component := templates.AddressTable(addresses)
	component.Render(c.Request.Context(), c.Writer)
}
