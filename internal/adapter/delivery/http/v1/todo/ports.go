package todo

import "github.com/nurlan42/todo/internal/domain/entity"

type Service interface {
	Create(task entity.TODO) error
	GetByID(ID string) (entity.TODO, error)
	GetAll() ([]entity.TODO, error)
	UpdateByID(ID string, todo entity.TODO) error
	DeleteByID(ID string) error
}
