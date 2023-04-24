package composites

import (
	"database/sql"
	handlerv1 "github.com/nurlan42/todo/internal/adapter/delivery/http/v1/todo"
	"github.com/nurlan42/todo/internal/adapter/repository/psql"
	"github.com/nurlan42/todo/internal/domain/service"
)

func NewTODO(db *sql.DB) *handlerv1.TODOHandler {
	repo := psql.NewTODO(db)
	serv := service.NewTODO(repo)
	handlerV1 := handlerv1.NewTODO(serv)
	return handlerV1
}
