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

