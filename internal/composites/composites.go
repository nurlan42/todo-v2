package composites

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurlan42/todo/config"
	"github.com/nurlan42/todo/internal/domain/service"
	"github.com/nurlan42/todo/internal/infrastructure/repository/postgres"
	"github.com/nurlan42/todo/internal/server"
	"github.com/nurlan42/todo/pkg/db"

	repoadapter "github.com/nurlan42/todo/internal/adapter/repository"
	servadapter "github.com/nurlan42/todo/internal/adapter/service"
	v1 "github.com/nurlan42/todo/internal/infrastructure/delivery/http/v1"
	swfiles "github.com/swaggo/files"
	ginsw "github.com/swaggo/gin-swagger"
)

func newTODO(gr *gin.RouterGroup, sqlDB *sql.DB) {
	repo := postgres.NewTODO(sqlDB)
	repoAdapter := repoadapter.New(repo)

	serv := service.NewTODO(repoAdapter)
	servAdapter := servadapter.New(serv)

	handler := v1.NewTODOHandler(servAdapter)
	handler.Init(gr)

	return
}

func NewApp(cfg *config.Config) error {
	sqlDB, err := db.Connect(cfg.DB.Psql)
	if err != nil {
		return err
	}

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/swagger/*any", ginsw.WrapHandler(swfiles.Handler))

	v1Handler := v1.NewHandler()

	v1Route := v1Handler.Init(router)

	newTODO(v1Route, sqlDB)

	if err := server.Run(cfg, router); err != nil {
		return err
	}

	return nil
}
