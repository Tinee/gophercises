// Package domain is the domain layer.
package domain

// Todo is the struct we're using when we want to represent
// something we're going to do in the future.
type Todo struct {
	ID      int    `json:"todoId"`
	Message string `json:"message"`
}

// BoltClient is the struct that keeps the underlaying bolt database pointer.
// It exposes a different services that we can use to interact with different structs.
type BoltClient interface {
	TodoService() TodoService
}

// TodoService have different methods that we could use to interact with Todos.
type TodoService interface {
	Create(Todo) error
	All() ([]Todo, error)
}
