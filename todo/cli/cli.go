package cli

import (
	"log"
	"strings"

	"github.com/Tinee/gophercises/todo/domain"
)

type Action struct {
	Kind string
	Todo domain.Todo
}

type Cli struct {
	s   domain.TodoService
	log *log.Logger
	Action
}

func NewCli(
	s domain.TodoService,
	l *log.Logger,
	a Action,
) *Cli {

	return &Cli{
		s,
		l,
		a,
	}
}

// Exectue exectues the action we provided.
func (c *Cli) Exectue(a Action) error {
	switch strings.ToLower(a.Kind) {
	case "list":
		todos, err := c.s.All()
		if err != nil {
			return domain.ErrGetAll
		}
		c.logTodoList(todos)
	case "create":

		err := c.s.Create(c.Todo)
		if err != nil {
			return domain.ErrCreate
		}
		c.log.Printf("Todo successfully created")
	case "do":
		err := c.s.Delete(c.Todo.ID)
		if err != nil {
			return domain.ErrCreate
		}

		c.log.Printf("Todo successfully removed")
	default:
		return domain.ErrInvalidAction
	}
	return domain.ErrInvalidAction
}

func (c *Cli) logTodoList(t []domain.Todo) {
	for _, v := range t {
		c.log.Printf("%v. %v\n", v.ID, v.Message)
	}
}
