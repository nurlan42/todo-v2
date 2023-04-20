package infrastructure

import (
	"github.com/nurlan42/todo/internal/infrastructure/delivery/http"
	"github.com/nurlan42/todo/internal/infrastructure/repository/postgres"
)

func NewTODORepo() *postgres.TODO {
	return postgres.NewTODO(nil, nil)
}

func NewTODOHandler() *http.Handler {
	return http.New(nil)
}
