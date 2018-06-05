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
		logger = log.New(os.Stdout, "", 0)
	)

	err = c.Open()
	defer c.Close()
	if err != nil {
		log.Fatalf("Error while opening a bolt connection: %v", err)
	}

	err = cli.NewCli(
		c.TodoService(),
		logger,
		cli.Action{
			Kind: action,
			Arg:  arg,
		},
	).Exectue()

	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
