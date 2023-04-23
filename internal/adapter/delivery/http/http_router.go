package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurlan42/todo/internal/domain"

	handlerv1 "github.com/nurlan42/todo/internal/adapter/delivery/http/v1"
	swfiles "github.com/swaggo/files"
	ginsw "github.com/swaggo/gin-swagger"
)

func newRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/swagger/*any", ginsw.WrapHandler(swfiles.Handler))
	return router

}

func NewHandler(services *domain.Services) *gin.Engine {
	router := newRouter()
	api := router.Group("api/")

	handlerV1 := handlerv1.New(services)
	handlerV1.Init(api)

	return router

}
