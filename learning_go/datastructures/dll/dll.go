package dll

/*
Element is the node doubly linked list. Each element has connection to next node and the previous node.
*/
type Element struct {
	prev, next *Element
	list       *CList // For operations like Next() from an element, e.next should be present and e.next should not be root. But element does not know about root. Thats the reason each element should know which list it belongs to
	Value      any
}

// CList Why should root be a sentinel node?. Can we not start from the node which is inserted
// Ans: We need the root element as it acts as a buffer nil node. See the init function for its use case
type CList struct {
	Root Element // root is a sentinel node (guard stands and keep watch). Modelled it as a value in api, convenient for users to interact.
	Len  int     // len is except the root element.
}

// lazyInit is used to initialise the list, in-case the user doesn't call the New function rather uses new(CList) and proceeds to using the dll methods.
func (l *CList) lazyInit() *CList {
	return Init(l)
}

// Init sets up root next and prev pointer to point to itself.
// Reason: When the first new element is inserted, the new element must point back the root and the root should point back to the new element.
func Init(l *CList) *CList {
	l.Root.next = &l.Root // wrap around the root.
	l.Root.prev = &l.Root
	return l
}

// New provides a pointer back to the user, the user mostly uses the methods of CList, so we haven't provided value type back.
func New() *CList {
	l := new(CList)
	return Init(l)
}

func (l *CList) PushFront(v any) {
	l.lazyInit()
	insertValue(v, &l.Root)
	l.Len++
}

func (l *CList) PushBack(v any) {
	l.lazyInit()
	insertValue(Element{Value: v}, l.Root.prev)
}

func (e *Element) Next() *Element {
	// e.next should be present and e.next should not be root. But element does not know about root. Thats the reason each element should know which list it belongs to
	if p := e.next; e.list != nil && p != &e.list.Root {
		return p
	}
	return nil
}

func insertValue(v any, at *Element) {
	insert(&Element{Value: v}, at)
}

// insert for PushFront use cases at is the root.
func insert(ne *Element, at *Element) *Element {

	// establish cnx from and to new element
	ne.next = at.next
	at.next.prev = ne

	// establish cnx from root
	at.next = ne
	ne.prev = at

	at.prev = ne.next // root back cnx
	return ne
}
