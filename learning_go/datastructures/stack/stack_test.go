package stack

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
)

func setup(size, rlimit int) *Stack[int] {
	s := New[int]()
	for range size {
		s.Push(rand.Intn(rlimit))
	}
	return s
}

func Test_stackCreation(t *testing.T) {
	data := []struct {
		name     string
		expected *Node[int]
		stack    *Stack[int]
	}{
		{
			name:     "empty stack creation",
			expected: nil,
			stack:    New[int](),
		},
		{
			name:     "create stack with multiple elements",
			expected: &Node[int]{val: 20, next: &Node[int]{val: 10}},
			stack: func() *Stack[int] {
				ms := New[int]()
				ms.Push(10)
				ms.Push(20)
				return ms
			}(),
		},
		{},
	}

	for _, dt := range data {
		t.Run(dt.name, func(t *testing.T) {
			if dt.stack.head != dt.expected {

			}
		})
	}
	s := New[int]()
	got := s.head

	if got != nil {
		t.Errorf("got: %v, expected: %v", got, nil)
	}
}

func Test_stackCreationWithHead(t *testing.T) {
	const tv = 10.0
	s := NewHead[float64](&Node[float64]{val: tv, next: nil})

	if s.head.val != 10.0 {
		t.Errorf("got: %v, expected: %v", s.head.val, tv)
	}
}

func Test_stackPushOps(t *testing.T) {
	const tv = 10
	s := New[int]()
	s.Push(tv)

	if tv != s.head.val {
		t.Errorf("got %d, expected %d", s.head.val, tv)
	}
}

func Test_stackPopOps(t *testing.T) {
	ms := setUp()

	_, err := ms.Pop()

	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func Test_stackIterator(t *testing.T) {
	ms := setUp()

	for v := range ms.Iter() {
		fmt.Printf("value: %d\t", v)
	}

	// breaking or returning false to stop iteration
	iterator := ms.Iter()

	iterator(func(value int) bool {
		if value > 10 {
			return false
		}
		return true
	})

	// use with for-range
	for v := range ms.Iter() {
		if v == 20 {
			break // we might not be able to yield a func with false, use break to quite execution.
		}
	}
}

func Test_stackPopSingleElement(t *testing.T) {
	ms := New[int]()
	ms.Push(10)
	v, err := ms.Pop()

	if err != nil {
		t.Errorf("error: %v", err)
	}

	if v != 10 {
		t.Errorf("got %d, expected %d", v, 10)
	}
}

func Test_emptyStackPopOp(t *testing.T) {
	s := New[int]()

	_, err := s.Pop()
	var got string
	if err != nil {
		got = err.Error()
	}
	me := errors.New("stack is empty")
	if me.Error() != got {
		t.Errorf("got %v, expected %v", err, me)
	}
}

func Test_stackLen(t *testing.T) {
	ms := setUp()

	got := ms.Len()
	expected := 10

	if got != expected {
		t.Errorf("got %d, expected %d", got, expected)
	}
}

func setUp() *Stack[int] {
	s := New[int]()
	return s
}

/*
* Benchmark Tests
`go test -v -bench=. -bechmem` (produces benchmark with memory allocation)
Example O/P: Benchmark_stackPushOps-12       21222189                53.81 ns/op           16 B/op          1 allocs/op
1st column denotes the name of the benchmark. 12 denotes the GOMAXPROCS value
2nd column denotes the number of iterations
3rd column denotes the time taken per operation
4th column denotes the memory allocated per operation
5th column denotes the number of allocations per operation
*/
func Benchmark_stackPushOps(b *testing.B) {
	s := New[int]()
	const rc = 20
	for i := 0; i < b.N; i++ {
		s.Push(rand.Intn(rc))
	}
}
