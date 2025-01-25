## Notes on Slices in Go

* A slice is an array with dynamic capacity, it is different from an array in go which has a fixed
  size.
* A slice typically accepts a length and a capacity.

## Dynamic Resizing in Go Internals

* Slice is a dynamically resizable array. An array is of fixed size and is used to store data of a prescribed type.
* Internally an array is a sequence of pointers. The pointer allocation pattern depends on the size of the data type.
* The allocation also depends on whether it can be stored in stack or in the heap which the go runtime decides basis the
  capacity rules.

Slice is represented as a SliceHeader which is pointer descriptor(variable) to the backing the array.

``` go
type SliceHeader struct {
  data    unsafe.Pointer
  cap     int
  len     int
}
```

To understand this we need to know what is unsafe.Pointer?

Properties of a Pointer different from other types

- A pointer value of any type can be converted to a pointer (Example: a := 10, b := &a, c := unsafe.Pointer(b))
- Pointer can be converted to a pointer value of any type.
- A uintptr can be converted to a Pointer
- A pointer can be converted to a uintptr

***Pointers allow to defeat the type system and read and write arbitrary memory***

Following patterns are valid, other patterns can be lead to unreliable or invalid results now or in the future.

## Conversion from *T1 to *T2

``` go
math.Float64bits

func Float64bits(f float64) uint64 {
  return *(*uint64(unsafe.Pointer(&f)))
}
```

## Conversion of pointer to uintptr & not vice-versa

***uintptr***

- Conversion of pointer to uintptr is usually used to get the memory address of value pointer at in integer format.
- It is usually used for printing.
- uintptr is just an integer not a reference. Even if the Pointer holds some reference to an object, converting to
  uintptr means uintptr is just an integer and does not point to any reference.
- It does not track the object being relocated in memory nor does it stop the gc from collecting the object.
- uintptr cannot be stored in a variable. (u := uintptr(p) //INVALID) before it converted to a pointer. (p =
  unsafe.Pointer(u + offset))

### Conversion of Pointer to uintptr and back, with arithmetic

- if p points to an allocated object, it can be advanced through the object by converting it into uintptr, addition of
  an offset and conversion back to Pointer.

``` go
p = unsafe.Pointer(uintptr(p) + offset)
```

- It is also valid to round pointers using &^.

### Common pattern is to access fields of a struct or elements of an array

f := unsafe.Pointer(&s.f)
f := unsafe.Pointer(uintptr(&s) + unsafe.Offsetof(&s.f)) // equivalent of 1

e := unsafe.Pointer(&x[i])
e := unsafe.Pointer(uintptr(unsafe.Pointer(&x[0])) + i*unsafe.Sizeof(x[0])) // get the data type size multiple by i to
move offset.

## Usage in SysCalls

- A SysCall function call passes the uintptr directly to the operating system and lets it manipulate it to a pointer
  depending on the calling function.
- `syscall.SysCall(SYS_READ, uintptr(fd), uintptr(unsafe.Pointer(p)), uintptr(n))`
- The compiler handles a Pointer converted to a uintptr in the argument list of
  a call to a function implemented in assembly by arranging that the referenced
  allocated object, if any, is retained and not moved until the call completes,
  even though from the types alone it would appear that the object is no longer
  needed during the call.

``` go
u := uintptr(unsafe.Pointer(p)) // INVALID. uintptr cannot be assigned to a variable
syscall.SysCall(SYS_READ, uintptr(fd), u, uintptr(n))
```

