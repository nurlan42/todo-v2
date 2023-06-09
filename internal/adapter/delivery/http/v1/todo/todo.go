package todo

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/nurlan42/todo/internal/domain/entity"

	log "github.com/sirupsen/logrus"
)

type Handler struct {
	service Service
	log     log.Entry
}

func New(s Service) *Handler {
	return &Handler{
		service: s,
	}
}

// Create godoc
// @Summary creates data in database
// @Produce json
// @Success 201
// @Router /api/v1/adapter [get]
func (h *Handler) Create(c *gin.Context) {
	logg := h.log.WithFields(log.Fields{"publicID": "4c120c2e-2c8b-4204-a54d-79c21d6f4b31"})

	logg.Info("start")

	var todo entity.TODO
	if err := c.BindJSON(&todo); err != nil {
		logg.Error(fmt.Errorf("BindJSON: %s; %s", err, debug.Stack()))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err := h.service.Create(todo)
	if err != nil {
		logg.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.Status(http.StatusCreated)
	logg.Info("end")

}

func (h *Handler) GetByID(c *gin.Context) {
	logg := h.log.WithFields(log.Fields{"userID": "userID"})

	id := c.Param("id")

	todo, err := h.service.GetByID(id)
	if err != nil {
		logg.Error(err)
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (h *Handler) GetAll(c *gin.Context) {
	logg := h.log.WithFields(log.Fields{"userID": "userID"})

	allTODO, err := h.service.GetAll()
	if err != nil {
		logg.Error(err)
		return
	}

	c.JSON(http.StatusOK, allTODO)
}

func (h *Handler) UpdateByID(c *gin.Context) {
	logg := h.log.WithFields(log.Fields{"userID": "userID"})

	todoID := c.Param("id")

	var todo entity.TODO
	if err := c.BindJSON(&todo); err != nil {
		logg.Error(err)
		return
	}

	err := h.service.UpdateByID(todoID, todo)
	if err != nil {
		logg.Error(err)
		return
	}

	c.Status(http.StatusOK)
}
func (h *Handler) DeleteByID(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteByID(id)
	if err != nil {
		//	adapter err
		return
	}

	c.Status(http.StatusOK)
}
