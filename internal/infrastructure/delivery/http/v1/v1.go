package v1

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

// NewHandler takes service
func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Init(r *gin.Engine) *gin.RouterGroup {
	return r.Group("/v1")
}
