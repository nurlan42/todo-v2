package composites

import (
	"database/sql"
	handler "github.com/nurlan42/todo/internal/adapter/delivery/http/v1/todo"
	repository "github.com/nurlan42/todo/internal/adapter/repository/psql/todo"
	service "github.com/nurlan42/todo/internal/domain/service/todo"
	"github.com/pressly/goose/v3"
	log "github.com/sirupsen/logrus"
)

func NewTODO(db *sql.DB) *handler.Handler {
	//embFs := migrations.New()

	//goose.SetBaseFS(embFs)

	if err := goose.SetDialect("postgres"); err != nil {
		log.Panic(err)
	}

	if err := goose.Down(db, "internal/migrations"); err != nil {
		log.Panic(err)
	}

	repo := repository.New(db)
	//repo := cache.NewTODO()
	serv := service.New(repo)
	handlerV1 := handler.New(serv)
	return handlerV1
}
