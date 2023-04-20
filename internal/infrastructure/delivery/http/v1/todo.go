package v1

import (
	"fmt"
	"github.com/nurlan42/todo/internal/domain/entity"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

// @title Blueprint Swagger API
// @version 1.0
// @description Swagger API for Golang Project Blueprint.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email martin7.heinz@gmail.com

// @BasePath /api/v1/adapter
func (h *TODOHandler) initTODO(gr *gin.RouterGroup) {
	td := gr.Group("/adapter")
	{
		td.POST("/", h.Create)
		td.GET("/:id", h.GetByID)
		td.GET("/", h.GetAll)
		td.PUT("/:id", h.UpdateByID)
		td.DELETE("/:id", h.DeleteByID)
	}
}

// Create godoc
// @Summary creates data in database
// @Produce json
// @Success 201
// @Router /api/v1/adapter [get]
func (h *TODOHandler) Create(c *gin.Context) {
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

func (h *TODOHandler) GetByID(c *gin.Context) {
	logg := h.log.WithFields(log.Fields{"userID": "userID"})

	id := c.Param("id")

	todo, err := h.service.GetByID(id)
	if err != nil {
		logg.Error(err)
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (h *TODOHandler) GetAll(c *gin.Context) {
	logg := h.log.WithFields(log.Fields{"userID": "userID"})

	allTODO, err := h.service.GetAll()
	if err != nil {
		logg.Error(err)
		return
	}

	c.JSON(http.StatusOK, allTODO)
}

func (h *TODOHandler) UpdateByID(c *gin.Context) {
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
func (h *TODOHandler) DeleteByID(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteByID(id)
	if err != nil {
		//	adapter err
		return
	}

	c.Status(http.StatusOK)
}
