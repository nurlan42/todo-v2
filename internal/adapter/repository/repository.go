package repository

import (
	"database/sql"
	"github.com/nurlan42/todo/internal/adapter/repository/psql"
	"github.com/nurlan42/todo/internal/domain/entity"
)

type TODORepos interface {
	Create(td entity.TODO) error
	GetByID(ID string) (entity.TODO, error)
	GetAll() ([]entity.TODO, error)
	UpdateByID(ID string, todo entity.TODO) error
	DeleteByID(ID string) error
}

type Repos struct {
	TODORepo TODORepos
}

func New(sqlDB *sql.DB) *Repos {
	return &Repos{
		TODORepo: psql.NewTODO(sqlDB),
	}
}
