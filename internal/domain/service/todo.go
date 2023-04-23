package service

import (
	"github.com/google/uuid"
	"github.com/nurlan42/todo/internal/adapter/repository"
	"github.com/nurlan42/todo/internal/domain/entity"
)

type TODO struct {
	//repo TODORepository
	repo *repository.Repos
}

func NewTODO(r *repository.Repos) *TODO {
	return &TODO{
		repo: r,
	}
}

func (t *TODO) Create(td entity.TODO) error {
	td.ID = uuid.New().String()
	return t.repo.TODORepo.Create(td)
}

func (t *TODO) GetByID(ID string) (entity.TODO, error) {
	return t.repo.TODORepo.GetByID(ID)
}

func (t *TODO) GetAll() ([]entity.TODO, error) {
	return t.repo.TODORepo.GetAll()
}

func (t *TODO) UpdateByID(ID string, todo entity.TODO) error {
	return t.repo.TODORepo.UpdateByID(ID, todo)
}

func (t *TODO) DeleteByID(ID string) error {
	return t.repo.TODORepo.DeleteByID(ID)
}
