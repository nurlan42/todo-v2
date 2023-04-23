package domain

import (
	"github.com/nurlan42/todo/internal/adapter/repository"
	"github.com/nurlan42/todo/internal/domain/entity"
	"github.com/nurlan42/todo/internal/domain/service"
)

type TODOService interface {
	Create(task entity.TODO) error
	GetByID(ID string) (entity.TODO, error)
	GetAll() ([]entity.TODO, error)
	UpdateByID(ID string, todo entity.TODO) error
	DeleteByID(ID string) error
}

type Services struct {
	TODOServ TODOService
}

func NewServices(r *repository.Repos) *Services {
	return &Services{
		TODOServ: service.NewTODO(r),
	}
}
