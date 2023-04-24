package todo

import "github.com/gin-gonic/gin"

// @title Blueprint Swagger API
// @version 1.0
// @description Swagger API for Golang Project Blueprint.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email martin7.heinz@gmail.com

// @BasePath /api/v1/todo1
func InitRoutes(r *gin.RouterGroup, h *TODOHandler) {
	td := r.Group("/todo")
	{
		td.POST("/", h.Create)
		td.GET("/:id", h.GetByID)
		td.GET("/", h.GetAll)
		td.PUT("/:id", h.UpdateByID)
		td.DELETE("/:id", h.DeleteByID)
	}
}
