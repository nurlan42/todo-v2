package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/nurlan42/todo/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	services *domain.Services
	log      log.Entry
}

func New(s *domain.Services) *Handler {
	return &Handler{
		services: s,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("v1/")
	{
		h.initTODO(v1)
	}
}
