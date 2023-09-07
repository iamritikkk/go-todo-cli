package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ritikrc/todo-cli"
)

const (
	todoFile = ".todo.json"
)

func main() {

	add := flag.Bool("add", false, "Add a new task")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		todos.Add("Radha Soami")
		err := todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	default:
		fmt.Fprintln(os.Stdout, "Invalid Command")
		os.Exit(0)
	}

}
