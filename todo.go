/* Keep track of items left in a project or activity.
The tool will save the list of items in a file using the JSON format.
Get input from standard input (STDIN) and command-line parameters.
Use environment variables to modify program work.
Display information back to the user through standard output (STDOUT)
and output errors with the standard error (STDERR) stream
for proper CLI error handling.*/

package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// private to this package
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

/*
	CREATES NEW TODO ITEM AND ADDS IT TO THE LIST OF ITEMS

receiver is a pointer to the type cuz Add method needs to modify
the content of the List by adding more items,
Otherwise, the method would change a copy of the list instead,
*/
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Now(),
	}
	// need to dereference the pointer to the List type with *l
	// in the append call to access the underlying slice.
	*l = append(*l, t)
}

/*
	MARKS ITEM AS COMPLETED

Complete method doesn’t modify the list, so it doesn’t require a pointer receiver
But it’s a good practice to keep the entire method set of a single type
with the same receiver type.
*/
func (l *List) Complete(i int) error {
	ls := *l
	// fmt.Println("completed ls List:", ls)
	// [{NewTask false 2023-02-23 14:33:07.481815288 +0300 +03 m=+0.000574216 2023-02-23 14:33:07.481815385 +0300 +03 m=+0.000574316}]
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d doesnt exist", i)
	}

	// adjust index for 0 based index
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()
	return nil
}

// DELETES TODO ITEM FROM THE LIST
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}

	// Adjust index for 0 based index​
	*l = append(ls[:i-1], ls[i:]...)

	return nil
}

// ENCODES LIST AS JSON AND SAVES IT
// TODO: ioutil is deprecated, rework that later
func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, js, 0644)
}

// OPEN FILE AND DECODES JSON INTO THE List
// handles file doesnt exist or empty
func (l *List) Get(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}

	// ok
	return json.Unmarshal(file, l)

}

// exposed lib to work with todo items
func main() {

}
