package todo

import (
	"fmt"
	"github.com/nurlan42/todo/internal/domain/entity"

	"github.com/google/uuid"
	"runtime/debug"
)

type TODO struct {
	cache map[string]entity.TODO
}

func New() *TODO {
	return &TODO{
		cache: make(map[string]entity.TODO),
	}
}

func (t *TODO) Create(td entity.TODO) error {
	td.ID = uuid.New().String()

	t.cache[td.ID] = td

	return nil
}

func (t *TODO) GetByID(ID string) (entity.TODO, error) {

	td, found := t.cache[ID]
	if !found {
		return entity.TODO{}, fmt.Errorf("GetByID: not found; %s", debug.Stack())
	}

	return td, nil
}

func (t *TODO) GetAll() ([]entity.TODO, error) {
	if len(t.cache) == 0 {
		return nil, fmt.Errorf("repository empty: %s", debug.Stack())
	}
	var all []entity.TODO
	for _, val := range t.cache {
		all = append(all, val)
	}
	return all, nil
}

func (t *TODO) UpdateByID(ID string, todo entity.TODO) error {
	if _, found := t.cache[ID]; !found {
		return fmt.Errorf("UpdateByID(%s) = not found; ID = %s; %s", ID, ID, debug.Stack())
	}

	t.cache[ID] = todo

	return nil
}

func (t *TODO) DeleteByID(ID string) error {
	delete(t.cache, ID)
	return nil
}
