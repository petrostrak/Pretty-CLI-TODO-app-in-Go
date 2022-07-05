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
	add := flag.Bool("add", false, "add a new task")
	complete := flag.Int("complete", 0, "mark a task as completed")
	del := flag.Int("del", 0, "delete a task")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(TODO_FILE); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	//go run ./cmd/todo -add
	case *add:
		todos.Add("sample todo")
		err := todos.Store(TODO_FILE)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	//go run ./cmd/todo -complete=1
	case *complete > 0:
		err := todos.Complete(*complete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		err = todos.Store(TODO_FILE)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *del > 0:
		err := todos.Delete(*del)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		err = todos.Store(TODO_FILE)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}
}
