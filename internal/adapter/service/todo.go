package service

import (
	"github.com/google/uuid"
	"github.com/nurlan42/todo/internal/domain/entity"
)

type Service interface {
	Create(task entity.TODO) error
	GetByID(ID string) (entity.TODO, error)
	GetAll() ([]entity.TODO, error)
	UpdateByID(ID string, todo entity.TODO) error
	DeleteByID(ID string) error
}

type TODO struct {
	Serv Service
}

func New(s Service) *TODO {
	return &TODO{
		Serv: s,
	}
}

func (t *TODO) Create(td entity.TODO) error {
	td.ID = uuid.New().String()
	return t.Serv.Create(td)
}

func (t *TODO) GetByID(ID string) (entity.TODO, error) {
	return t.Serv.GetByID(ID)
}

func (t *TODO) GetAll() ([]entity.TODO, error) {
	return t.Serv.GetAll()
}

func (t *TODO) UpdateByID(ID string, todo entity.TODO) error {
	return t.Serv.UpdateByID(ID, todo)
}

func (t *TODO) DeleteByID(ID string) error {
	return t.Serv.DeleteByID(ID)
}
