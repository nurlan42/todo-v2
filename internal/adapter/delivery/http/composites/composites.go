package composites

import (
	"database/sql"

	"github.com/nurlan42/todo/internal/adapter/repository/cache"

	handler "github.com/nurlan42/todo/internal/adapter/delivery/http/v1/todo"
	service "github.com/nurlan42/todo/internal/domain/service/todo"
)

func NewTODO(db *sql.DB) *handler.Handler {
	//repo := psql.NewTODO(db)
	repo := cache.NewTODO()
	serv := service.New(repo)
	handlerV1 := handler.New(serv)
	return handlerV1
}
