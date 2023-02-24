// CLI INTERFACE TO API
// 1) when without any arguments - will list the available to-do items.
// 2) when with one or more arguments, the command will concatenate
// the arguments as a new item and add it to the list.

// update - use `flag` to handle args options
// will accept:
// -list: Boolean  When used, will list all todo items.
// -task: string  When used, will include the string arg as a new todo item in the list.
// -complete: integer . When used, will mark the item number as completed.

package main

import (
	"bufio" //read data from the STDIN input stream
	"flag"
	"fmt"
	"io" //use the io.Reader interface
	"os"

	"strings" // join command-line arguments to compose a task name
	"todo"
)

// if env var is empty
var todoFileName = ".todo.json"

// getTask decides where to get the description for a new task from:
// args or STDIN
// [variadic function](https://go.dev/ref/spec#Function_types)
// if any arguments were provided as the parameter args.
// if yes, returns all of them concatenated with a " "
// if no - uses the bufio.Scanner to scan for a single input line
// on the provided io.Reader interface.
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
		return "", fmt.Errorf("task cannot be empty")
	}

	return s.Text(), nil
}

func main() {

	// assigned vars are pointers, need to be dereferenced by *
	task := flag.String("task", "", "task to be included in list")
	add := flag.Bool("add", false, "add task to list")
	list := flag.Bool("list", false, "list all tasks")
	complete := flag.Int("complete", 0, "Item to be complete")

	flag.Parse()

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage of %s info: ", os.Args[0])
		flag.PrintDefaults()
	}

	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	// fmt.Println("task", task) //task 0xc0000962c0

	// extract the address of an empty instance of todo.List.
	// This variable represents the todo items list to use throughout the code.
	l := &todo.List{}
	// fmt.Println(l) // &[]

	// read items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// fmt.Println(l.Get(todoFileName)) // nil at first

	// do based on the flag provided
	switch {
	// case len(os.Args) == 1: // no args, filename only
	case *list:
		// 1) list items
		// for _, item := range *l {

		// 	// show only pending tasks
		// 	if !item.Done {
		// 		fmt.Println(item.Task)
		// 	}
		// }

		// 2) satisfying Stringer interface provide the type to any
		// formatting function that expects a string
		// can call the fmt.Print function, which requires no format specifier,
		// as the format comes from the Stringer interface
		// implemented by the l var of type todo.List.
		fmt.Print(l) // l is *todo.List

	// -------- complete flag
	case *complete > 0: // default
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// save new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *task != "": // ad new task
		l.Add(*task)
		// and save
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	// when any args (excluding flags) are provided
	// they will be used as new task
	case *add:
		// can use the os.Stdin as the 1st param cuz
		// its type *os.File implements the io.Reader interface.
		// flag.Args returns all the remaining non-flag arguments
		// use ... to expand the slice into a list of vals as expected by the func.
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		l.Add(t)

		// save
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	default:
		// any flag provided
		// prompt := "proper use:\n -task [task name] add new task to task list,\n -list list all items,\n -complete [number] complete chosen task,\n "
		fmt.Fprintln(os.Stderr, "invalid flag")
		os.Exit(1)
	}

}
