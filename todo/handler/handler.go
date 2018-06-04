package handler

import (
	"errors"
	"log"
	"strings"

	"github.com/Tinee/gophercises/todo/domain"
)

var (
	// ErrNoOp happeneds when the caller pass down an incorrect argument.
	ErrNoOp   = errors.New("it seems that the argument you passed wrong")
	ErrGetAll = errors.New("something went wrong when getting the todos")
)

type action int

const (
	List action = iota
	Add
	Do
	NoOp
)

type Handler struct {
	s      domain.TodoService
	action action
	log    *log.Logger
}

func NewHandler(
	action string,
	s domain.TodoService,
	output *log.Logger) *Handler {

	return &Handler{
		action: figureOutAction(action),
		s:      s,
	}
}

// Exectue exectues the action we provided.
func (h *Handler) Exectue() error {
	switch h.action {
	case List:
		todos, err := h.s.All()
		if err != nil {
			return ErrGetAll
		}
		handleList(h.log, todos)
	case NoOp:
		return ErrNoOp

	}
	return ErrNoOp
}

func handleList(w *log.Logger, t []domain.Todo) {
	for i, v := range t {
		log.Printf("%v. %v\n", i, v.Message)
	}
}

func figureOutAction(action string) action {
	switch strings.ToLower(action) {
	case "list":
		return List
	case "do":
		return Do
	case "add":
		return Add
	default:
		return NoOp
	}
}
