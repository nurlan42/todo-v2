package http

import "github.com/gin-gonic/gin"

func New() *gin.Engine {
	router := gin.New()

	router.Use(AuthMiddleware)

	return router
}
