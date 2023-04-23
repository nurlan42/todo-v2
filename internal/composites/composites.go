package composites

import (
	"github.com/nurlan42/todo/config"
	"github.com/nurlan42/todo/internal/server"
	"github.com/nurlan42/todo/pkg/db"

	httprouter "github.com/nurlan42/todo/internal/adapter/delivery/http"
	handltodo "github.com/nurlan42/todo/internal/adapter/delivery/http/v1/todo"
	repotodo "github.com/nurlan42/todo/internal/adapter/repository/psql/todo"
	servtodo "github.com/nurlan42/todo/internal/domain/service/todo"
)

func New(cfg *config.Config) error {
	sqlDB, err := db.Connect(cfg.DB.Psql)
	defer sqlDB.Close()

	if err != nil {
		return err
	}

	router := httprouter.New()

	v1 := router.Group("/v1")

	repo := repotodo.New(sqlDB)
	serv := servtodo.New(repo)
	handlerV1 := handltodo.New(serv)
	handlerV1.Init(v1)

	if err := server.Run(cfg.HTTP, router); err != nil {
		return err
	}

	return nil
}
