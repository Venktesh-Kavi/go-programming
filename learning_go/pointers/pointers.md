## Notes on Pointers

* A pointer is a variable whose values points to the address of another variable.

``` go
a := 10
b := &a
   a     b
|-----|------|
| 10  |  &a  |
|-----|------|

value    0   0   0   10
memory   1   2   3   4
variable a

value    0   0   0   1
memory   5   6   7   8
variable b

// Pointers have a default memory size in bytes
```

* In Go everything is pass by value.
* & is the address operator and it precedes a value type. Returns the address of the value type.
* `*` indirection operator. Precedes a variable of pointer type. It returns the value being pointed
  to. This is called dereferencing.
* Pointer type is type that represents a pointer. It represented a * before the type name.
* `new` creates a pointer variable. It returns a zero values instance of the provided type.
* Use `new` sparingly. For structs use &Struct{}
* When pointers are passed to functions, a copy of the pointer is sent.

``` go
// here a copy of pointer b is passed. Both the original and the copy point to the same value.
func(b *int) {
     
}
```

``` go
func MakeFoo(foo *Foo) error {
  foo.Name = "dupefoo"
  foo.CreatedAt = time.Now()
}
//instead
func MakeFoo() (Foo, error) {
  return foo{
    Name: "dupefoo",
    CreatedAt: time.Now(),
  }, nil
}
```

* Use pointer parameters to modify a variable only when the parameter expects an interface. Eg..
  json.Unmarshall(b []bytes, data interface{})

### Pointer passing performance

* If a struct is large enough there is a performance benefit of passing structs as pointers in
  either the parameters or while returning.
* Time to pass a pointer to a function is almost constant for all data sizes, roughly 1
  nanosecond. (As size of pointer is same for all datatypes uint)
* Passing a value to a function takes about a millisecond for data of size 10MB.
* Returning a pointer type for a data point less than a MB is ***slower*** than returning a value.
  For a datastructure of size 100 bytes, returning a value takes 10ns, whereas a pointer might take
  30ns.
* This is flipped when the data point size is larger than an MB.

### Points on Maps and Slices

* Any modifications made to the map by passing it as parameter to a function, results in the
  original map to change. The reason itself it map is implemented as a pointer to a struct. We are
  passing a copy of the pointer hence the underlying struct changes if any modifications are done.
* **Do not use maps as input parameters or return values for any language, especially on public
  APIs (APIs refers to method/functions which users interact in a library or your code).
    * Maps are a bad choice because they say nothing about the keys or value, one has to trace
      through the code to understand.
    * From the stand of immutability, maps are bad because the only way to know what ended up in the
      map, is to trace through all functions which interacted with the map.
    * This prevents API design to be self-documented.
* Rather pass structs.

* Slices behaviour is slightly different. When a slice is passed as a parameter to a function, (
  internally slice is represented as a struct of len and capacity int fields and a pointer to the
  backing array), a copy of the underlying struct is passed.
* Any modifications to the slice in the function is reflected, as the pointer to the backing array
  is modified. This is not the case with len and capacity. This is not the case with len and
  capacity.
* So appends are not visible to the slice outside the function.
* So as a paradigm assume that a slice is not modified by a function. The function should document
  if it modifies the slices.

### Slices as Buffers

Typically programming languages code is written like below:

``` java
try(BufferedReader br = new BufferedReader(new FileReader("file.txt"))) {
  String line = br.readLine();
  List<Bytes> lineBytes = line.getBytes();
  process(lineBytes);
} catch(IOException ex) {
  // some ex
}

/**
 In the above code everytime there is a new allocation of list of bytes. In a garbage collected languages, the collector takes care of removing them, but its redundant work.
**/
```

``` go
// use slices as reusable buffers
file, err = os.Open("file.txt")
if err != nil {
  return err
}

defer file.Close()

data := make([]byte, 100)

for {
  count, err := file.Read(data)
  if err != nil {
    return err
  }
  if count == 0 {
    return nil
  }
  
  process(data[:count])
}

/**
In the above code, we use slice as a buffer to read data from the file source.
We can't change the len and capacity of the slice while passing it to a function, but we can modify the contents.
we create a slice of 100 bytes and each time we loop through we copy the next block of 100 bytes into the slice.
We pass on the populated contents of the buffer for processing.
**/
```

* The above pattern in go is a very efficient and idiomatic way to write go code with efficient
  memory utilisation.