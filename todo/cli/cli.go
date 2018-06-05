package cli

import (
	"log"
	"strconv"
	"strings"

	"github.com/Tinee/gophercises/todo/domain"
)

type Action struct {
	Kind string
	Arg  string
}

type Cli struct {
	s   domain.TodoService
	log *log.Logger
	Action
}

func New(
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
func (c *Cli) Exectue() error {
	switch strings.ToLower(c.Kind) {
	case "list":
		return c.handleList()
	case "create":
		return c.handleCreate()
	case "do":
		return c.handleDo()
	default:
		return domain.ErrInvalidAction
	}
}

func (c *Cli) handleList() error {
	todos, err := c.s.All()
	if err != nil {
		return domain.ErrGetAll
	}

	for _, v := range todos {
		c.log.Printf("%v. %v\n", v.ID, v.Message)
	}

	return nil
}

func (c *Cli) handleDo() error {
	id, err := strconv.Atoi(c.Arg)
	if err != nil {
		c.log.Println("Error: invalid argument")
		return err
	}

	err = c.s.Delete(id)
	if err != nil {
		return domain.ErrDelete
	}

	c.log.Printf("Successfully deleted the todo with id %v", id)
	return nil
}

func (c *Cli) handleCreate() error {
	err := c.s.Create(domain.Todo{
		Message: c.Arg,
	})

	if err != nil {
		return domain.ErrCreate
	}

	c.log.Printf("Successfully added the todo \"%v\"", c.Arg)

	return nil
}
