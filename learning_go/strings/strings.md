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