package tree

import (
	"fmt"
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

	//fmt.Printf("address of root: %p\n", &root)
	root.Clear() // implicitly pointer of root is passed

	got := root.Height()
	var expected *Node[int]
	expected = &Node[int]{
		Value: 0,
		Left:  nil,
		Right: nil,
	}
	if *root != *expected {
		t.Fatalf("got %v, expected: %v", root, expected)
	}
	if got != 1 {
		t.Fatalf("default height of tree is 1 because of zero value, got: %d, expected: %d\n", got, 1)
	}
}

func TestNilPointerTree(t *testing.T) {
	var enode *Node[int] // this pointer points to nil

	got := enode.Height()
	expected := 0
	if got != expected {
		t.Fatalf("got %d, expected %d", got, expected)
	}
}
