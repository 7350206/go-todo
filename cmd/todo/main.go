// CLI INTERFACE TO API
// 1) when without any arguments - will list the available to-do items.
// 2) when with one or more arguments, the command will concatenate
// the arguments as a new item and add it to the list.

package main

import (
	"fmt"
	"os"
	"strings"
	"todo"
)

// hardcode for now
const todoFileName = ".todo.json"

func main() {
	// extract the address of an empty instance of todo.List.
	// This variable represents the to-do items list you’ll use throughout the code.
	l := &todo.List{}
	// fmt.Println(l)

	// read items from file
	if err := l.Get(todoFileName); err != nil {
		// use the standard error (STDERR) output instead of the
		// standard output (STDOUT) to display error messages
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// what the program should do based on the arguments it received
	switch {
	case len(os.Args) == 1: // no args, filename only
		// list current items
		for _, item := range *l {
			fmt.Println(item.Task)
		}

	// add new task as args sum,
	// no check for separate args, at all for now
	default:
		// concatenate all args with a space
		item := strings.Join(os.Args[1:], " ")

		l.Add(item) // add composed item

		// try to save
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	}

}
