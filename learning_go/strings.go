package main;

import "fmt"

/**

  Notes on strings in go

  * Strings are immutable in Go
  * Variables assigned to a string are descriptors pointing to a memory block. The descriptors also carry additional info like len along with them
    * programming languages like C have \0 to represent the end of string. To get a length of string, users might have to diregard the \0. Go provides it inplace for us.
  * Len operation on a string, provides the total byte size of the utf8 enchoded unicode string. (Typical general purpose strings use ASCII (american english literals)). Unicode corresponds to international chars (e' french literal etc..,)
    `* utf8 encoding is used to reduce the size of the unicode literals which typically higher number of bytes. eg, the e'lite is represented as [195, 142, 105, 98, 120, 112]. Even though elite has only 5 characters the byte representation takes 6 as e' takes 2 bytes
  

**/
func main() {
  s := "hello world"
  
}
