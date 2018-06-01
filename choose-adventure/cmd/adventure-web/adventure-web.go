package main

import (
	"fmt"
	"os"

	"github.com/Tinee/gophercises/choose-adventure/adventure"
)

func main() {
	f, err := os.Open("./gopher.json")
	if err != nil {
		fmt.Printf("Error while reading file: %v", err)
		os.Exit(1)
	}

	a := adventures.New(f)

	fmt.Printf("%+v", a)
}
