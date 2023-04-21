package repository

import (
	"github.com/google/uuid"
	"github.com/nurlan42/todo/internal/domain/entity"
)

type Repository interface {
	Create(task entity.TODO) error
	GetByID(ID string) (entity.TODO, error)
	GetAll() ([]entity.TODO, error)
	UpdateByID(ID string, todo entity.TODO) error
	DeleteByID(ID string) error
}

type TODO struct {
	repo Repository
}

func New(r Repository) *TODO {
	return &TODO{
		repo: r,
	}
}

func (t *TODO) Create(td entity.TODO) error {
	td.ID = uuid.New().String()
	return t.repo.Create(td)
}

func (t *TODO) GetByID(ID string) (entity.TODO, error) {
	return t.repo.GetByID(ID)
}

func (t *TODO) GetAll() ([]entity.TODO, error) {
	return t.repo.GetAll()
}

func (t *TODO) UpdateByID(ID string, todo entity.TODO) error {
	return t.repo.UpdateByID(ID, todo)
}

func (t *TODO) DeleteByID(ID string) error {
	return t.repo.DeleteByID(ID)
}
