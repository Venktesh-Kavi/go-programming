## Range Over Function Types

[Ref Go Docs](https://go.dev/blog/range-functions)
Talk in **GopherCon 2024**

* The blog discusses on a Set data type, whose internal representation is a map.
* The author takes a union of two sets example to walk through the need for range over function
  types.

``` go
type Set Struct {
    m    map[E]struct{}
}
func NewSet[E comparable]() *Set[E] {
    return &Set[E]{make(map[E]struct{})}
}
func Union(E comparable](s1, s2 *Set[E]) *Set[E] {
    result := NewSet[E]()
    for e := range s1 {
        result.Add(e)
    }
    for e := range s2 {
        result.Add(e)
    }
    return result
}
```

* In the above example we want to iterate over the internal maps to get the union of two sets.
* The map internal type is unexported, making it unavailable for other packages to access them (
  though it is accessible within the set package).
* The author proposes a couple of ways to solve this:

### Sol 1: Push Set Elements

> Expose a Push function to accept a function as an argument to indicate till when to push the
> elements.

``` go
func (s *Set[E]) Push[E](f func(E) bool) {
    for _, e := range s.m {
        fmt.Println("pushing value", e)
        if !f(e) {
            return
        }
    } 
}
```