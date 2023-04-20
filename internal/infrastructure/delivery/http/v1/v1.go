package v1

import (
	"github.com/gin-gonic/gin"
	slog "github.com/sirupsen/logrus"
)

type TODOHandler struct {
	service TODOService
	log     *slog.Logger
}

// NewHandler takes service
func NewHandler(u TODOService) *TODOHandler {
	return &TODOHandler{
		service: u,
		log:     slog.New(),
	}
}

func (h *TODOHandler) Init(gr *gin.RouterGroup) {
	v1 := gr.Group("/v1")
	{
		h.initTODO(v1)
	}
}
