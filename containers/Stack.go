package containers

// TODO: decide if the embedded struct implementation is useful... The list semantics are still available
type Stack[T any] struct {
	DoubleLinkedList[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(p T) {
	if s == nil {
		return
	}
	s.InsertAtHead(p)
}

func (s *Stack[T]) Pop() (T, bool) {
	var ret T
	if s == nil {
		return ret, false
	}
	ret, ok := s.Get(s.DoubleLinkedList.Length() - 1)
	if !ok {
		return ret, false //not found
	}
	s.DeleteAtTail() // remove it from the list
	return ret, true
}
