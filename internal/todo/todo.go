package todo

import (
	"context"
	"sync"
)

type ToDo struct {
	id          int
	description string
}

type Service struct {
	todos   map[int]ToDo
	nextID  int
	todosMu *sync.Mutex
}

func NewService() (*Service, error) {
	todos := make(map[int]ToDo)
	return &Service{
		todos:   todos,
		nextID:  0,
		todosMu: &sync.Mutex{},
	}, nil
}

func (svc *Service) CreateTODO(ctx context.Context, todo string) (int, error) {
	svc.todosMu.Lock()
	defer svc.todosMu.Unlock()

	t := ToDo{
		id:          svc.nextID,
		description: todo,
	}

	svc.todos[svc.nextID] = t
	return t.id, nil
}

func (svc *Service) GetTODO(ctx context.Context, id int) (ToDo, error) {
	svc.todosMu.Lock()
	defer svc.todosMu.Unlock()

	todo, _ := svc.todos[id]

	return todo, nil
}
