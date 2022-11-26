package main

import (
	"cl_apps/todo"
	"flag"
	"fmt"
	"os"
)

const todoFileName = ".todo.json"

func main() {
	// Parsing command line flags
	task := flag.String("task", "", "Task to be included in the ToDo list")
	desc := flag.String("desc", "", "Description of the task to be added")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	flag.Parse()

	l := &todo.TodoList{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	switch {
	// For no extra arguments, print the list
	case *list:
		// List current to do items
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	case *complete > 0: 
	// complete the given item 
	if err := l.Complete(*complete); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Save the new list
	if err := l.Save(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	case *task != "":
		// Add the task
		l.AddTask(*task, *desc)
		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
	}
	default:
		// Invalid flag provided
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}

}
