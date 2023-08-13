package containers

// TODO: decide if the embedded struct implementation is useful... The list semantics are still available
type Queue[T any] struct {
	DoubleLinkedList[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(p T) {
	if q == nil {
		return
	}
	q.InsertAtHead(p)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var ret T
	if q == nil {
		return ret, false
	}
	ret, ok := q.Get(q.DoubleLinkedList.Length() - 1)
	if !ok {
		return ret, false //not found
	}
	q.DeleteAtTail() // remove it from the list
	return ret, true
}
