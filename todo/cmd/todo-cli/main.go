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
		c      = bolt.NewClient("todo.db")
		logger = log.New(os.Stdout, "", 0)
	)

	defer c.Close()
	err = c.Open()
	if err != nil {
		log.Fatalf("Error while opening a bolt connection: %v", err)
	}
	svc := c.TodoService()

	h := cli.NewCli(svc, logger)

	err = h.Exectue(action)
}
