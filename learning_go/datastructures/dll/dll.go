package dll

/*
Element is the node doubly linked list. Each element has connection to next node and the previous node.
*/
type Element struct {
	prev, next *Element
	Value      any
}

// CList Why should root be a sentinel node?. Can we not start from the node which is inserted.
type CList struct {
	Root Element // root is a sentinel node (guard stands and keep watch). Modelled it as a value in api, convenient for users to interact.
	Len  int     // len is except the root element.
}

// lazyInit is used to initialise the list, in-case the user doesn't call the New function rather uses new(CList) and proceeds to using the dll methods.
func (l *CList) lazyInit() *CList {
	return Init(l)
}

// Init sets up root next and prev pointer to point to itself.
func Init(l *CList) *CList {
	l.Root.next = &l.Root // wrap around the root.
	l.Root.prev = &l.Root
	return l
}

func New() *CList {
	l := new(CList)
	return Init(l)
}

func (e *Element) PushFront(v any) {
	insertValue(v, l)
}

func insertValue(v any, e *Element) {
	insert(&Element{Value: v}, e)
}

func insert(ne *Element, ee *Element) {
}
