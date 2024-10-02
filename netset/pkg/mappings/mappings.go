package mappings

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/konstfish/whitelist-portal/netset/pkg/controllers"
)

var Router *gin.Engine

func CreateUrlMappings() {
	Router = gin.Default()

	Router.Use(controllers.Cors())

	// internals
	Router.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// api
	v1 := Router.Group("/api/v1")
	{
		v1.GET("/list", controllers.GetAllIPAddresses)
		v1.GET("/list/:user", controllers.GetUserIPAddresses)
	}
}
