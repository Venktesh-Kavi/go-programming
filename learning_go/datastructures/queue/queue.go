package queue

import (
	"cmp"
	"container/list"
)

const DEFAULT_CAP = 10

type Queue[T cmp.Ordered] struct {
	data [DEFAULT_CAP]*T
	sp   int
	ep   int
}

func (q *Queue[T]) Push(v T) {
	list.New()
	q.data[q.ep] = &v
	q.ep++
	if q.ep == DEFAULT_CAP {
		q.extendCapacity()
	}
}

func (q *Queue[T]) extendCapacity() {
	// how to extend capacity, without declaring it as slice. As it just allocate infinite memory again basis some default capacity logic.
	// If capacity is hard coded, how to extend the capacity as the type enforces a fixed array.

	// copy the contents of data from src to dest
	nq := Queue[T]{data: [DEFAULT_CAP]*T{}, sp: q.sp}
	for _, d := range q.data {
		nq.data[nq.ep] = d
		nq.ep++
	}
	nq.sp = q.sp
}

func (q *Queue[T]) Pop() *T {
	if q.sp == 0 { // skip
	}
	v := q.data[q.sp]
	q.data[q.sp] = nil // garbage collection to kick in
	q.sp++
	return v
}

func (q *Queue[T]) len() int {
	return q.ep - q.sp
}
