package composites

import (
	"database/sql"

	"github.com/nurlan42/todo/internal/adapter/repository/cache"

	handlertodo "github.com/nurlan42/todo/internal/adapter/delivery/http/v1/todo"
	servtodo "github.com/nurlan42/todo/internal/domain/service/todo"
)

func NewTODO(db *sql.DB) *handlertodo.Handler {
	//repo := psql.NewTODO(db)
	repo := cache.NewTODO()
	serv := servtodo.New(repo)
	handlerV1 := handlertodo.NewTODO(serv)
	return handlerV1
}
