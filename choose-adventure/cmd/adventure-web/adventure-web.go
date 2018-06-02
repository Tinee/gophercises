package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Tinee/gophercises/choose-adventure/adventure"
)

func main() {
	mux := http.NewServeMux()
	f, err := os.Open("./gopher.json")

	if err != nil {
		fmt.Printf("Error while reading file: %v", err)
		os.Exit(1)
	}

	as, err := adventures.New(f, "./adventure.html")

	if err != nil {
		fmt.Printf("Error creating adventures: %v", err)
		os.Exit(1)
	}

	// adventures.Find()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		k := r.URL.Query().Get("adventure")
		as.WriteByKey(w, k)
	})

	http.ListenAndServe(":3000", mux)
}
