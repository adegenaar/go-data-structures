// todo: fill in the package documentation
package containers

import "fmt"

// TODO: document the singleNode information
type singleNode[T any] struct {
	data T
	next *singleNode[T]
}

func newSingleNode[T any](value T) *singleNode[T] {
	return &singleNode[T]{
		data: value,
		next: nil,
	}
}

type SingleLinkedList[T any] struct {
	head   *singleNode[T]
	tail   *singleNode[T]
	length int
}

func NewSingleLinkedList[T any]() *SingleLinkedList[T] {
	return &SingleLinkedList[T]{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (l *SingleLinkedList[T]) InsertAtHead(value T) {
	newNode := newSingleNode(value)
	newNode.next = l.head
	l.head = newNode

	//special case, inserting into an empty list
	if l.length == 0 {
		l.tail = newNode
	}
	l.length++
}

// Add a node of type T value at the end of the linked list
func (l *SingleLinkedList[T]) InsertAtTail(value T) {
	// special case of the empty list
	if l.head == nil {
		l.InsertAtHead(value)
		return
	}

	newNode := newSingleNode(value)
	l.tail.next = newNode
	l.tail = newNode
	l.length++
}

// Add a node of type T value before the indexth node in the linked list.
// If index equals the length of the linked list, the node will be appended to the end of the linked list.
// If index is greater than the length, the node will not be inserted.
func (l *SingleLinkedList[T]) InsertAt(index int, value T) error {
	// special case: input index is out of range, return error
	size := l.length
	if index < 0 || index > size {
		return fmt.Errorf("index out of range")
	}

	// special case: index is 0, insert at head by use InsertAthead method
	if index == 0 {
		l.InsertAtHead(value)
		return nil
	}

	// special case: index equals the length of list, insert at tail by use InsertAtTail method
	if index == size {
		l.InsertAtTail(value)
		return nil
	}

	// general case: find the node before the index, then insert a new node after the found node
	cur := l.head
	for i := 0; i < index-1; i++ {
		cur = cur.next
	}
	newNode := newSingleNode(value)
	newNode.next = cur.next
	cur.next = newNode
	l.length++

	return nil
}

func (l *SingleLinkedList[T]) DeleteAtHead() error {
	if l.head == nil {
		return fmt.Errorf("list is empty")
	}

	l.head = l.head.next
	l.length--

	// special case, deleted the last node
	if l.head == nil {
		l.tail = nil
	}
	return nil
}

func (l *SingleLinkedList[T]) DeleteAtTail() error {
	if l.head == nil {
		return fmt.Errorf("list is empty")
	}
	// special case: the list has only one node
	if l.head.next == nil {
		return l.DeleteAtHead()
	}

	// general case: find the second last node of list, then delete its' next node
	cur := l.head
	for cur.next != l.tail {
		cur = cur.next
	}
	cur.next = nil

	l.tail = cur
	l.length--
	return nil
}

// Delete the indexth node in the linked list, if the index is valid.
func (l *SingleLinkedList[T]) DeleteAt(index int) error {
	size := l.length
	// special case: input index is out of range, return error
	if index < 0 || index > size-1 {
		return fmt.Errorf("index out of range")
	}
	// special case: index is 0, delete at head by use DeleteAthead method
	if index == 0 {
		return l.DeleteAtHead()
	}

	// general case: find the node before the index, then delete its' next node
	cur := l.head
	for i := 0; i < index-1; i++ {
		cur = cur.next
	}
	cur.next = cur.next.next
	l.length--
	return nil
}

// Get the value of the indexth node in the linked list, if the index is valid.
func (l *SingleLinkedList[T]) Get(index int) (T, bool) {
	// special case: input index is out of range, return error
	var t T
	if index < 0 || index > l.length-1 {
		return t, false
	}
	// general case: find the node at the index, then return its' value
	cur := l.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.data, true
}

func (l *SingleLinkedList[T]) Length() int {
	return l.length
}

//
func (l *SingleLinkedList[T]) Values() []T {
	values := []T{}
	cur := l.head
	for cur != nil {
		values = append(values, cur.data)
		cur = cur.next
	}
	return values
}
