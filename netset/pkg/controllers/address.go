package controllers

import (
	"errors"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/konstfish/whitelist-portal/netset/pkg/cache"
)

func writeHeader(c *gin.Context) {
	c.Header("Content-Type", "text/plain")

	currentTime := time.Now().UTC()
	c.Writer.Write([]byte("#\n# whitelist-portal netset\n"))
	c.Writer.Write([]byte("# " + currentTime.Format(time.RFC3339) + "\n#\n"))
}

func GetAllIPAddresses(c *gin.Context) {
	ipList, err := cache.GetAllIPAddresses(c.Request.Context())
	if err != nil {
		log.Println(err)
		c.AbortWithError(500, errors.New("unable to retrieve"))
		return
	}

	writeHeader(c)
	for _, ip := range ipList {
		c.Writer.Write([]byte(ip + "\n"))
	}
}

func GetUserIPAddresses(c *gin.Context) {
	user := c.Param("user")

	ipList, err := cache.GetUserIPAddresses(c.Request.Context(), user)
	if err != nil {
		log.Println(err)
		c.AbortWithError(500, errors.New("unable to retrieve"))
		return
	}

	writeHeader(c)
	for _, ip := range ipList {
		c.Writer.Write([]byte(ip + "\n"))
	}
}
