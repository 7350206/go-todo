### improve list output format
- if api not owned : the only way is cli implementation
- on owned api:  try go interface to inplement output formatting directly

interface in Go implements a contract but, unlike other languages, Go interfaces define `only behavior and not state`.  
This means that an interface defines `what a type should do` and not `what type of data it should hold`.

- to satisfy an interface, a type needs only to `implement all the methods defined in the interface` with the same signature.  
- satisfying an interface doesn’t require explicit declaration. Types will implicitly implement an interface by `implementing all the defined methods`.

By implicitly satisfying an interface, a given type can be used anywhere that interface is expected, enabling code decoupling and reuse.

---
implement the `Stringer` interface on the `todo.List` type in api  
The Stringer interface is defined in the `fmt` package

```go
type Stringer interface {
  String() string 
}
```

Any types that implement the method `String, which returns a string`, satisfy the `Stringer` interface. By satisfying this interface, you can provide the type to any formatting function that expects a string.


### easy file save
inprove flexibility with env vars

- another flag can be added to choise file to save/
- use env var for that (avoid typing any time)

`os` package provides functions to handle both the environment and env vars.
will use `os.Getenv("TODO_FILENAME")` to retrieve the value of the environment var identified by the name `TODO_FILENAME.`

`to check later: export​​ ​​TODO_FILENAME=new-todo.json`
`go run ./main.go -task "check env"`
`cat new-todo.json`

### Capturing Input from STDIN
common way cli interact with one another is by accepting input from the STDIN stream.  
need the ability to add new tasks via STDIN, allowing users to pipe new tasks from other command-line tools.

new helper function `getTask` will determine where to get the input task from  
- leverages Go interfaces again by accepting the `io.Reader` interface as input.  
it’s a good practice to `take interfaces as function arguments` instead of concrete types.  
This approach increases the flexibility of functions by allowing `different types to be used as input` as long as they satisfy the given interface.  
As interfaces are implicitly satisfied, it’s `common in Go to have simple interfaces composed of one or two methods.` The io.Reader is an example of a simple interface that provides a lot of flexibility.

use this interface whenever expect to read data  
types such as `files`, `buffers`, `archives`, `HTTP requests`, and others satisfy this interface.

using it, decouple implementation from specific types, allowing code to work with any types that implement the `io.Reader` interface.

TODO:
- make the flag `-del` to delete an item from the list. Use the Delete method from the API to perform the action.

- Add another flag to enable verbose output, showing information like date/time.

- Add another flag to prevent displaying completed items.

- Update the custom usage function to include additional instructions on how to provide new tasks to the tool.

- Include test cases for the remaining options, such as -complete.

- Update the tests to use the TODO_FILENAME environment variable instead of hard-coding the test file name so that it doesn’t cause conflicts with an existing file.

- Update the `getTask` function allowing it to handle multiline input from STDIN. Each line should be a new task in the list.

:)