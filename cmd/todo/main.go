package main

import (
	"bufio"
	"cl_apps/todo"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

const todoFileName = ".todo.json"

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
		}
	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
		}
	if len(s.Text()) == 0 {
		return "", fmt.Errorf("task cannot be blank")
	}
	return s.Text(), nil
}

func main() {
	// Parsing command line flags
	desc := flag.String("desc", "", "Description of the task to be added")
	add := flag.Bool("add", false, "Add task to the ToDo list")
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

	case *add:
	// When any arguments (excluding flags) are provided, they will be
	// used as the new task
	t, err := getTask(os.Stdin, flag.Args()...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	l.AddTask(t, *desc)
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
