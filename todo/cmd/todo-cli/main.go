package main

import (
	"flag"
	"log"
	"os"

	"github.com/Tinee/gophercises/todo/cli"

	"github.com/Tinee/gophercises/todo/bolt"
)

func main() {
	flag.Parse()
	var (
		err    error
		action = flag.Arg(0)
		arg    = flag.Arg(1)
		c      = bolt.NewClient("todo.db")
	)

	err = c.Open()
	defer c.Close()
	if err != nil {
		log.Fatalf("Error while opening a bolt connection: %v", err)
	}

	err = cli.New(
		os.Stdout,
		cli.Action{
			Kind: action,
			Arg:  arg,
		},
		c.TodoService(),
	).Exectue()

	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
