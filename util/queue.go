package util

type Queue[T interface{}] struct {
	elems []T
}

func (q *Queue[T]) Offer(elem T) {
	if q.elems == nil {
		q.elems = make([]T, 0)
	}
	q.elems = append(q.elems, elem)
}

func (q *Queue[T]) Poll() T {
	var a T
	if q.IsEmpty() {
		return a
	}
	elem := q.elems[0]
	q.elems = q.elems[1:]
	return elem
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.elems) == 0
}
