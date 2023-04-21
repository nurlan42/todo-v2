package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	swfiles "github.com/swaggo/files"
	ginsw "github.com/swaggo/gin-swagger"
)

func New() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/swagger/*any", ginsw.WrapHandler(swfiles.Handler))
	return router

}
