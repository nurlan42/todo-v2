package domain

import "github.com/nurlan42/todo/internal/domain/service"

func NewTODOServ() *service.TODO {
	return service.NewTODO(nil)
}
