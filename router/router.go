package router

import (
	"github.com/gin-gonic/gin"
	"url-shortener/handler"
)

func Wrapper() *gin.Engine {

	router := gin.Default()
	_ = router.SetTrustedProxies([]string{"localhost:8080"})

	router.POST("/create", handler.CreateUrlHandler)
	router.GET("/:key", handler.GetUrlHandler)

	return router
}
