package v1

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/nurlan42/todo/internal/adapter/composites"
	"github.com/nurlan42/todo/internal/adapter/delivery/http/v1/todo"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHTTPV1Routes(db *sql.DB, router *gin.Engine) {
	routerV1 := router.Group("/v1")
	{
		{
			routerV1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
		todo.InitRoutes(routerV1, composites.NewTODO(db))

	}

}
