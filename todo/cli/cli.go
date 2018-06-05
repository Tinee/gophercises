package cli

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/Tinee/gophercises/todo/domain"
)

type Action struct {
	Kind string
	Arg  string
}

type Cli struct {
	s domain.TodoService
	w io.Writer
	Action
}

func New(w io.Writer, a Action, s domain.TodoService) *Cli {
	if w == nil {
		w = os.Stdout
	}
	return &Cli{s, w, a}
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
		fmt.Fprintf(c.w, "%v. %v\n", v.ID, v.Message)
	}

	return nil
}

func (c *Cli) handleDo() error {
	id, err := strconv.Atoi(c.Arg)
	if err != nil {
		fmt.Fprintln(c.w, "Error: invalid argument")
		return err
	}

	err = c.s.Delete(id)
	if err != nil {
		return domain.ErrDelete
	}

	fmt.Fprintf(c.w, "Successfully completed the todo with id %v", id)
	return nil
}

func (c *Cli) handleCreate() error {
	err := c.s.Create(domain.Todo{
		Message: c.Arg,
	})

	if err != nil {
		return domain.ErrCreate
	}

	fmt.Fprintf(c.w, "Successfully added the todo \"%v\"", c.Arg)
	return nil
}
