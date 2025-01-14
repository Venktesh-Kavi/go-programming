## Notes on Go Strings

* Strings are immutable in Go
* Variables assigned to strings are called descriptors, they hold
    * reference to the backing byte array
    * length information of the array
* Strings with the same literals may share the same backing array. This is called interning

### Len in Strings

* Len operation on a string, provides the total byte size of the utf8 encoded unicode string. (
  Typical general purpose strings use ASCII (american english literals)). Unicode corresponds to
  international chars (e' french literal etc..,)
* utf8 encoding is used to reduce the size of the unicode literals which typically higher number of
  bytes. eg, the e'lite is represented as [195, 142, 105, 98, 120, 112]. Even though elite has only
  5 characters the byte representation takes 6 as e' takes 2 bytes
* The backing array is an array of bytes in unicode
### Sample

``` go
s := "foo"
ns := "bar"

s = ns // s starts pointing to ns's backing byte arrary)

ns += "k" // ns now starts pointing to a new backing array "bark"

fmt.Println(s) // s still points to bar (backing array)
fmt.Println(ns) // ns is bark now
```

Slices of Strings

s := "Hello World!"

s[1:] => ello World
s[:1] => H

Strings are **Immutable** in Go. Reason
    - Thread Safety: strings can be shared between two go routines without any explicit locking.
    - Simplified logic: No need to worry about content changes in a string affecting other references.
    - Efficiency: Allocating new strings is added cost, go efficient memory allocator minimizes this cost.

If we want to just manipulate a string prefer bytes. 

``` go

s := "hello"
s1 := "world"

bytes := []byte(s)

// replace all l's to !.

for i, b := range bytes {
    if b == 'l' {
        bytes[i] = '!'
    }
}

return string(bytes)

```

If the given string has unicode characters use runes. As using bytes might split apart a multi byte character like a smiley.

### General Rule of Thumb:

#### Performance-Sensitive Code:

Converting a string to []rune incurs a cost because it involves decoding UTF-8 into Unicode code points.
For plain ASCII strings or situations where you only need byte-level operations, avoid []rune.

#### When Working with ASCII-Only Strings:

If the string is guaranteed to contain only single-byte ASCII characters, there’s no need to use []rune. Byte-level operations are sufficient and faster.

#### When You Don’t Need Character-Level Precision:

For tasks like finding substrings, comparing strings, or simple slicing in ASCII strings, stick to direct string operations.

