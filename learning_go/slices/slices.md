# Notes on Slices in Go

- A slice in Go is a dynamically resizable array, unlike an array which has a fixed size.
- Typically, a slice accepts a length and a capacity.

## Dynamic Resizing in Go Internals

- A slice is a dynamically resizable array, while an array is of fixed size and used to store data of a prescribed type.
- Internally, an array is a sequence of pointers. The pointer allocation pattern depends on the size of the data type.
- The allocation decision (whether to store in the stack or heap) is determined by the Go runtime based on capacity rules.
- A slice is represented as a `SliceHeader`, which is a descriptor pointing to the backing array.

```go
type SliceHeader struct {
  data unsafe.Pointer
  cap  int
  len  int
}
```

## Understanding `unsafe.Pointer`

### Properties of a Pointer

- A pointer value of any type can be converted to another pointer type.
  - Example:

    ```go
    a := 10
    b := &a
    c := unsafe.Pointer(b)
    ```

- A `uintptr` can be converted to a `Pointer`.
- A `Pointer` can be converted to a `uintptr`.
- Pointers allow bypassing the type system to read and write arbitrary memory.

**Note**: Following valid patterns can help avoid unreliable or invalid results now or in the future.

### Common Pointer Conversion Pattern

#### Conversion from T1 to T2

Example `math.Float64bits`:

```go
func Float64bits(f float64) uint64 {
  return *(*uint64)(unsafe.Pointer(&f))
}
```

#### Conversion of Pointer to uintptr & Not Vice-Versa

##### `uintptr`

- Conversion of a pointer to a `uintptr` is often used to get the memory address of the value pointed to in an integer format.
- Typically used for printing.
- `uintptr` is just an integer, not a reference. It does not track object relocation nor prevent garbage collection.
- `uintptr` cannot be directly assigned to a variable before being converted back to a pointer.

Example of converting pointer to `uintptr` and back, with arithmetic:

```go
p = unsafe.Pointer(uintptr(p) + offset)
```

- Rounding pointers using `&^` is also valid.

##### Accessing Fields of a Struct or Elements of an Array

Example patterns:

```go
f := unsafe.Pointer(&s.f)
f := unsafe.Pointer(uintptr(&s) + unsafe.Offsetof(s.f)) // equivalent to the above
e := unsafe.Pointer(&x[i])
e := unsafe.Pointer(uintptr(unsafe.Pointer(&x[0])) + i*unsafe.Sizeof(x[0])) // get data type size, multiplied by i to move offset
```

## Usage in SysCalls

- A syscall function call passes the `uintptr` directly to the operating system and lets it manipulate it to a pointer depending on the calling function.
- Example:

  ```go
  syscall.Syscall(SYS_READ, uintptr(fd), uintptr(unsafe.Pointer(p)), uintptr(n))
  ```

- The compiler ensures that a pointer converted to a `uintptr` in an argument list of a call to an assembly function is retained and not moved until the call completes.

## Removing an Element from a Slice

- Given an index, there are two ways to remove an element from a slice:

```go
// RemoveElement
func RemoveElement(nums []int, idx int) []int {
    nums = append(nums[:idx], nums[idx+1:]...)
    return nums
}

func RemoveEleByShifting(nums []int, idx int) []int {
    copy(nums[idx:], nums[idx+1:])
    return nums[:len(nums)-1]
}
```

- In the first approach `append` can possibly create new slice header with a new len. Though the backing array is the
  same. The returned slice can be a new view of the same backing array.

```go
aux := []int{1, 2, 3, 4, 5}
nAux := append(aux[:idx], aux[idx+1:]...)
// nAux and aux will be different slice descriptors.
fmt.Println(&aux, &nAux)
```

Correct Approaches for Removing Element

Ref: https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
- Use the slices.Delete method stable available from go 1.21
- Or do not touch the initial slice at all. Example
 
```go
func RemoveElement(nums []int, idx int) []int {
	nn := make([]int, len(nums) - 1)
	i, k := 0
    for i<len(nn) {
        if i != idx {
            nn[k] = nums[i]
            k++
        } else {
			k++
        }
        i++
    }
}
```