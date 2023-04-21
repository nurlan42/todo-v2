package composites

import (
	"github.com/nurlan42/todo/config"
	"github.com/nurlan42/todo/internal/adapter/repository/psql"
	"github.com/nurlan42/todo/internal/domain/service"
	"github.com/nurlan42/todo/internal/server"
	"github.com/nurlan42/todo/pkg/db"

	httprouter "github.com/nurlan42/todo/internal/adapter/delivery/http"
	handlerv1 "github.com/nurlan42/todo/internal/adapter/delivery/http/v1"
)

func New(cfg *config.Config) error {
	sqlDB, err := db.Connect(cfg.DB.Psql)
	defer sqlDB.Close()

	if err != nil {
		return err
	}

	router := httprouter.New()

	v1 := router.Group("/v1")

	repo := psql.NewTODO(sqlDB)
	serv := service.NewTODO(repo)
	handlerV1 := handlerv1.NewTODO(serv)
	handlerV1.Init(v1)

	if err := server.Run(cfg.HTTP, router); err != nil {
		return err
	}

	return nil
}
