package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Tinee/gophercises/todo/handler"

	"github.com/Tinee/gophercises/todo/bolt"
)

func main() {
	flag.Parse()
	var (
		err    error
		action = flag.Arg(0)
		c      = bolt.NewClient("todo.db")
		logger = log.New(os.Stdout, "", log.LstdFlags)
	)
	fmt.Println(action)

	defer c.Close()
	err = c.Open()
	if err != nil {
		log.Fatalf("Error while opening a bolt connection: %v", err)
	}
	svc := c.TodoService()

	h := handler.NewHandler(action, svc, logger)
	err = h.Exectue()
	if err != nil {

	}
}
