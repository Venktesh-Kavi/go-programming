## Notes on Runes in Go

* Runes and Bytes are not the same
* A byte is a unicode representation of a character in go.
* A rune is a int32 representation of a unicode point.
* Use len(rune) to get the exact length of the string

``` go
s := "foo"

bytes := []bytes(s)
runes := []runes(s)
fmt.Printf("Bytes %v: , len: %d\n", bytes, len(bytes))
fmt.Printf("Runes %v: , len: %d\n", runes, len(runes))

// Note: ASCII - (0 - 127)

o/p: Bytes: [93, 116, 116], len: 3
o/p: Runes: [93, 116, 116], len: 3

ns := "foo 世界"
nb := []bytes(ns)
nr := []runes(ns)

o/p: Bytes: [93, 116, 116, 12, 128, 123, 231, 123, 123], len: 9
o/p: Runes: [93, 116, 116, 12, 689, 4021], len: 6
```