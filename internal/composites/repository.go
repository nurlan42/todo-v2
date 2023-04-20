package composites

import (
	"github.com/nurlan42/todo/internal/adapter/repository/todo"
	"github.com/nurlan42/todo/internal/infrastructure"
)

func NewTODORepo() *todo.TODO {
	repo := infrastructure.NewTODORepo()
	return todo.New(repo)
}

func NewTODOServ() {

}
