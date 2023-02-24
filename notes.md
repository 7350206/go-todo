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
