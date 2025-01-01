package mappings

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/konstfish/whitelist-portal/portal/pkg/controllers"
	"github.com/konstfish/whitelist-portal/portal/pkg/templates"
)

var Router *gin.Engine

func CreateUrlMappings() {
	Router = gin.Default()

	Router.Use(controllers.Cors())

	// static files (https://github.com/gin-gonic/gin/issues/2654)
	Router.StaticFileFS("static/main.css", "static/main.css", http.FS(templates.StaticFiles))
	Router.StaticFileFS("static/ip.js", "static/ip.js", http.FS(templates.StaticFiles))

	Router.GET("/", controllers.ValidateAuthHeader, controllers.GetIndex)

	/*Router.StaticFS("/static", http.FS(templates.GetStaticFiles()))
	Router.GET("", func(c *gin.Context) {
		c.FileFromFS("index.htm", http.FS(templates.GetStaticFiles()))
	})*/

	// internals
	Router.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// api
	v1 := Router.Group("/hx/v1")
	{
		v1.GET("/addresses", controllers.ValidateAuthHeader, controllers.GetAddressList)
		v1.POST("/addresses", controllers.ValidateAuthHeader, controllers.AddAddress, controllers.GetAddressList)
		v1.DELETE("/addresses/:address", controllers.ValidateAuthHeader, controllers.DeleteAddress, controllers.GetAddressList)
	}
}
