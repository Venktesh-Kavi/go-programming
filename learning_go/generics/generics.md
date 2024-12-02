## Notes on Generics in Go

Introduced in Go 1.18

1. Type parameters for functions and types
2. Types sets are defined by interfaces
3. Type inference

Type Parameters are defined as:
[P, Q constraint, R constraint]

``` go
fun min[T constraint.Ordered](x, y float64) float64 {
    if x < y {
        return x
    } else {
        return y
    }
}

m := min[float64] (2.56, 3.45)

// other way, assign a generic type to a non geneirc one
m := min[float64]
nm := min(2.56, 3.45)
```

Instantiation

* Compiler first substitutes the type arguments with the type parameters.
* Checks that the type arguments satisfy the constraints.

### Type Sets

* Interfaces define method sets.
* In the below example types Bar and Baz implement the interface Foo.

``` go
type Foo interface {
    A()
}

type Bar struct {
}

type Baz struct {
}

func (b Bar) A() {
    fmt.Println("A")
}

func (b Baz) A() {
    fmt.Println("A")
}
```

* Interfaces can also be used to define type sets. Since any type can implement the interface (here
  Foo). The interface can define what types can implement it

``` go
// Exmaple
interface {
    int|bool|float64
}

// constraint.Ordered
package constraints
type Ordered interface {
    Integer|Float|~string
}

// Here Integer and Float are interfaces which define type sets themselves.
// Typically we dont care about a specific type like int or string, rather we care about all the int/string types.
// ~ denotes all types whose underlying type is a string.

```

* Constraints can be declared inline like below:

``` go
[S interface{~[]E}]
// S is a type parameter that requires a type of anything whose underlying type is a slice.

[E interface{}]
// The E type here is not constrained by anything. So an empty interface.

Syntactic sugar was introduced to remove the interface word by doing the follow

[S ~[]E] and [E interface{}]
[S ~[]E] and [E any]

// introduced a new predeclared identifier `any`, alias for empty interface
const any = interface{}
```

### When to use generics?

* Write code, don't design types
    * Start by writing functions, add type parameters later.


* Use Type parameters in functions that work on slices, maps, channels of any element type. If a
  function code has parameters with the mentioned types and the function code doesn't make any
  assumption about the element types, then its ok to use generics/type parameters.
    * Eg.., function that returns a slice of keys of any map type.
* General purpose data structure. (Binary trees or linked list)
    * Using typed parameters in-place of interface types can often permit the data to be stored more
      efficiently.
    * Code can avoid type assertions and code can be full type checked at compile time.
    * When operating on type parameters, prefer functions over
      method [TODO: REVISIT] [Ref Time: ](https://youtu.be/Pa_e9EeCdy8)
    * We could have defined the Tree type with an E constrained by comparable or some type sets
      which have less implemented. (As there is a comparison operation with 0 in the findVal method
      in BT)
        * This means that the user of the tree type, must write a owning type and implement their
          own compare method to instantiate a tree
        * If a take a func as we have in Tree (cmp func) method, it makes it easier to pass the
          function. If the type already has a compare method, we could just pass
          `elementType.Compare()`
        * It easier to convert method to a function than adding a method to a data type.
        * ***Hence for general purpose data types prefer a function rather than writing a constraint
          that requires a method.***

``` go
// example binary tree datastructure

type Tree[T any] struct {
    cmp    func(T, T) int
    root    *node[T]
}

type node[T any] struct {
    left, right    *node[T]
    data           T
}

```

* When method looks the same for all types. (Example is sorting a slice) (When different types want
  to implement a method and the implementation of the different types look the same)
    * Irrespective of any slice the sort func is the same. We need to implement the len, swap and
      less methods.

``` go
// a generic slice type that implements the sort interface. For any slice the len and swap methods are the same.
type SliceFn[T any] struct {
    s   []T
    cmp func(T, T) bool
}

// sort interface methods
func (s SliceFn[T]) Len() { return len(s.s) }
func (s SliceFn[T]) Swap(i, j int) { return s.s[i], s.s[j] = s.s[j], s.s[i] }
func (s SliceFn[T]) Less(i, j int) { return s.cmp(s[i], s[j]) }


// usage of sortFn to sort a sliceFn
func SortFn[T any](s []T, cmp func(T, T) bool) {
    sort.Sort(SliceFn[T]{s, cmp})
}
// the slices types are identical and it ok to use generics here.
```

### When not to use generics?

* When just calling a method on the type argument.
    * Eg.., the `io.Reader`. Interface types provide a way for generic programming.
    * io.Reader provides a generic way to read data from file or some method which produces some
      data
      like random no generator.
    * If the only job is to call a method on a value data. Use the interface type parameter rather
      than
      a generic type parameter.
* When the implementation of a common method is different for each type. (reading from a file is
  different from reading from random no. generator, 2 different read method neither require a type
  parameter)
* When the operation is different for each type, even without a method. (Use reflection here).
    * Eg.. encoding/json package, we dont want every type we require support a marshal json method,
      can't use interface types. Also encoding a int type is nothing like encoding a struct type (so
      we cant use type parameters). Instead the package uses reflection

``` go
// OK
func SomeFn(r io.Reader) ([]byte, error) {
}

// Not OK
func SomeFn[T any](r T) ([]byte, error) {
}
// this make the type to implement the reader interface.

Why 1st is prefered over 2nd?
- Omitting the type parameter here, makes the function easier to write and read.
```


