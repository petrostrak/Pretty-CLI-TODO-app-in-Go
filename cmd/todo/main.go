package main

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/petrostrak/Pretty-CLI-TODO-app-in-Go"
)

const (
	TODO_FILE = ".todos.json"
)

func main() {
	add := flag.Bool("add", false, "add a new todo")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(TODO_FILE); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		todos.Add("sample todo")
		err := todos.Store(TODO_FILE)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}
}
