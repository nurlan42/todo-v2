package composites

import (
	"github.com/nurlan42/todo/config"
	"github.com/nurlan42/todo/internal/adapter/repository"
	"github.com/nurlan42/todo/internal/domain"
	"github.com/nurlan42/todo/internal/server"
	"github.com/nurlan42/todo/pkg/db"

	httprouter "github.com/nurlan42/todo/internal/adapter/delivery/http"
)

func New(cfg *config.Config) error {
	sqlDB, err := db.Connect(cfg.DB.Psql)
	defer sqlDB.Close()

	if err != nil {
		return err
	}

	repo := repository.New(sqlDB)
	serv := domain.NewServices(repo)
	router := httprouter.NewHandler(serv)

	if err := server.Run(cfg.HTTP, router); err != nil {
		return err
	}

	return nil
}
