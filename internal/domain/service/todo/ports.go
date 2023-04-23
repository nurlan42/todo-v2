package todo

import "github.com/nurlan42/todo/internal/domain/entity"

type Repository interface {
	Create(td entity.TODO) error
	GetByID(ID string) (entity.TODO, error)
	GetAll() ([]entity.TODO, error)
	UpdateByID(ID string, todo entity.TODO) error
	DeleteByID(ID string) error
}
