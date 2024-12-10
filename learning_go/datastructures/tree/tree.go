package tree

type Node[T comparable] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

func (n *Node[T]) InsertLeft(value T) *Node[T] {
	if n == nil {
		return createRootNode(value)
	}
	n.Left = &Node[T]{
		Value: value,
	}
	return n.Left
}

func (n *Node[T]) InsertRight(value T) *Node[T] {
	if n == nil {
		return createRootNode(value)
	}
	n.Right = &Node[T]{
		Value: value,
	}
	return n.Right
}

func createRootNode[T comparable](value T) *Node[T] {
	return &Node[T]{
		Value: value,
	}
}

// Height is the no of edges from a node to the leaf node.
func (n *Node[T]) Height() int {
	if n == nil {
		return 0
	}
	lh := n.Left.Height()
	rh := n.Right.Height()
	var res int
	if lh < rh {
		res = rh
	} else {
		res = lh
	}
	return 1 + int(res)
}

// Clear method receiver receives a copy of the pointer. performing n = nil, only affects the local copy.
func (n *Node[T]) Clear() {
	if n == nil {
		return
	}

	//fmt.Printf("address of receiver: %p\n", &n)
	if n.Left != nil {
		n.Left.Clear()
		n.Left = nil
	}

	if n.Right != nil {
		n.Right.Clear()
		n.Right = nil
	}
	var zero T
	n.Value = zero
}
