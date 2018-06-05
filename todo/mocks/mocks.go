package mock

import "github.com/Tinee/gophercises/todo/domain"

// TodoService is a mock for TodoService
type TodoService struct {
	CreateTodoFn      func(t domain.Todo) error
	CreateTodoInvoked bool

	DeleteTodoFn      func(id int) error
	DeleteTodoInvoked bool

	AllFn      func() ([]domain.Todo, error)
	AllInvoked bool
}

// Create mocks the underlaying Create on TodoService
func (s TodoService) Create(t domain.Todo) error {
	s.CreateTodoInvoked = true
	return s.CreateTodoFn(t)
}

// Delete mocks the underlaying Delete on TodoService
func (s TodoService) Delete(id int) error {
	s.CreateTodoInvoked = true
	return s.DeleteTodoFn(id)
}

// All mocks the underlaying All on TodoService
func (s TodoService) All() ([]domain.Todo, error) {
	s.AllInvoked = true
	return s.AllFn()
}
