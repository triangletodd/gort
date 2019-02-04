package server

import (
	"github.com/gin-gonic/gin"
	"github.com/triangletodd/gort/internal/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	url := new(controllers.UrlController)
	health := new(controllers.HealthController)

	router.GET("/service/status", health.Status)

	router.GET("/url", url.RetrieveUrls)
	router.POST("/url", url.CreateUrl)
	router.GET("/url/:short", url.RetrieveUrl)

	// This is hacky, but necessary to make /:name redirect properly.
	//   https://github.com/gin-gonic/gin/issues/388
	router.NoRoute(url.Handler)

	return router
}
