package tree

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_TreeCreation(t *testing.T) {
	root := &Node[int]{
		Value: 10,
	}

	// root is a nil pointer, test passing a nil pointer to insert method.
	/**
		 10
		/	\
		20	30
		/
	   40
	*/
	ln1 := root.InsertLeft(20)
	root.InsertRight(30)

	ln1.InsertLeft(40)

	got := root.Height()
	fmt.Printf("Height of tree: %d\n", got)
	expected := 3

	if got != expected {
		t.Fatalf("got %d, expected %d\n", got, expected)
	}
}

func TestClearTree(t *testing.T) {
	root := &Node[int]{
		Value: 10,
	}
	root.InsertLeft(20)
	rn1 := root.InsertRight(30)
	rn1.InsertRight(40)

	fmt.Printf("root pointer: %v\n", reflect.TypeOf(root).Kind())
	root.Clear() // implicitly pointer of root is passed

	got := root.Height()
	var expected int
	if got != expected {
		t.Fatalf("got %d, expected %d\n", got, expected)
	}
}
