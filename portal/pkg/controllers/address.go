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

func AddAddress(c *gin.Context) {
	user := userFromContext(c)

	var address cache.AddressListEntry
	if err := c.Bind(&address); err != nil {
		c.AbortWithStatus(400)
		return
	}

	// check limit
	if limit, err := cache.GetAddressCount(c.Request.Context(), user); err != nil {
		c.AbortWithStatus(500)
		return
	} else if int(limit) >= UserAddressLimit {
		out, _ := templates.ErrorMessage("Max address limit reached").Render()
		c.Writer.Write([]byte(out))
		return
	}

	// validate form input
	if valid, reason := validateAddress(address); !valid {
		out, _ := templates.ErrorMessage(reason).Render()
		c.Writer.Write([]byte(out))
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

	for _, item := range addresses {
		out, _ := templates.AddressTableEntry(item).Render()

		c.Writer.Write([]byte(out))
	}

	if len(addresses) == 0 {
		out, _ := templates.ErrorMessage("No addresses yet!").Render()
		c.Writer.Write([]byte(out))
	}
}
