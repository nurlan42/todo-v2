package http

import (
	"github.com/nurlan42/todo/internal/infrastructure/delivery/http/v1"
	"net/http"

	"github.com/gin-gonic/gin"
	swfiles "github.com/swaggo/files"
	ginsw "github.com/swaggo/gin-swagger"
)

type Handler struct {
	serv v1.TODOService
}

// New takes service adapter
func New(s v1.TODOService) *Handler {
	return &Handler{
		serv: s,
	}
}

// @title Blueprint Swagger API
// @version 1.0
// @description Swagger API for Golang Project Blueprint.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email martin7.heinz@gmail.com

// @BasePath /api/v1/adapter
func (h *Handler) Init() *gin.Engine {

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/swagger/*any", ginsw.WrapHandler(swfiles.Handler))

	h.initEndpoints(router)

	return router
}

func (h *Handler) initEndpoints(r *gin.Engine) {
	api := r.Group("/api")

	v1Handler := v1.NewHandler(h.serv)

	v1Handler.Init(api)
}
