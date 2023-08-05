// todo: fill in the package documentation
package godatastructures

import "fmt"

// TODO: document the doubleNode information
type doubleNode[T any] struct {
	data T
	next *doubleNode[T]
	prev *doubleNode[T]
}

func newDoubleNode[T any](value T) *doubleNode[T] {
	return &doubleNode[T]{
		data: value,
		next: nil,
		prev: nil,
	}
}

type DoubleLinkedList[T any] struct {
	head   *doubleNode[T]
	tail   *doubleNode[T]
	length int
}

func NewDoubleLinkedList[T any]() *DoubleLinkedList[T] {
	return &DoubleLinkedList[T]{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (l *DoubleLinkedList[T]) InsertAtHead(value T) {
	newNode := newDoubleNode(value)
	newNode.next = l.head

	if l.head != nil {
		l.head.prev = newNode
	}
	// newNode.prev is nil when inserting at the head of the list
	l.head = newNode

	//special case, inserting into an empty list
	if l.length == 0 {
		l.tail = newNode
	}
	l.length++
}

// Add a node of type T value at the end of the linked list
func (l *DoubleLinkedList[T]) InsertAtTail(value T) {
	// special case of the empty list
	if l.head == nil {
		l.InsertAtHead(value)
		return
	}

	newNode := newDoubleNode(value)
	l.tail.next = newNode
	newNode.prev = l.tail
	l.tail = newNode
	l.length++
}

// Add a node of type T value before the indexth node in the linked list.
// If index equals the length of the linked list, the node will be appended to the end of the linked list.
// If index is greater than the length, the node will not be inserted.
func (l *DoubleLinkedList[T]) InsertAt(index int, value T) error {
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
	newNode := newDoubleNode(value)
	newNode.prev = cur
	newNode.next = cur.next
	cur.next = newNode
	l.length++

	return nil
}

func (l *DoubleLinkedList[T]) DeleteAtHead() error {
	if l.head == nil {
		return fmt.Errorf("list is empty")
	}

	l.head = l.head.next
	l.length--

	// special case, deleted the last node
	if l.head == nil {
		l.tail = nil
	} else {
		l.head.prev = nil
	}
	return nil
}

func (l *DoubleLinkedList[T]) DeleteAtTail() error {
	if l.head == nil {
		return fmt.Errorf("list is empty")
	}
	// special case: the list has only one node
	if l.head.next == nil {
		return l.DeleteAtHead()
	}

	// general case: find the second last node of list, then delete its' next node
	l.tail = l.tail.prev
	l.tail.next = nil

	l.length--
	return nil
}

// Delete the indexth node in the linked list, if the index is valid.
func (l *DoubleLinkedList[T]) DeleteAt(index int) error {
	size := l.length
	// special case: input index is out of range, return error
	if index < 0 || index > size-1 {
		return fmt.Errorf("index out of range")
	}
	// special case: index is 0, delete at head by use DeleteAthead method
	if index == 0 {
		return l.DeleteAtHead()
	}
	// special case: index is equal to size -1 (i.e. the tail), delete at tail by use DeleteAtTail method
	if index == size-1 {
		return l.DeleteAtTail()
	}
	// general case: find the node before the index,
	//               then delete its' next node
	// optimization: if the index is small, start at the head otherwise start at the tail
	if index < size/2 {
		cur := l.head
		for i := 0; i < index-1; i++ {
			cur = cur.next
		}
		cur.next = cur.next.next
		cur.next.prev = cur
	} else {
		cur := l.tail
		for i := size - 1; i > index; i-- {
			cur = cur.prev
		}

		cur.prev.next = cur.next
		cur.next.prev = cur.prev
	}
	l.length--
	return nil
}

// Get the value of the indexth node in the linked list, if the index is valid.
func (l *DoubleLinkedList[T]) Get(index int) (T, bool) {
	// special case: input index is out of range, return error
	var t T
	size := l.length
	if index < 0 || index > size-1 {
		return t, false
	}
	// general case: find the node at the index, then return its' value
	// optimization: if the index is small, start at the head otherwise start at the tail
	if index < size/2 {
		cur := l.head
		for i := 0; i < index; i++ {
			cur = cur.next
		}
		return cur.data, true
	} else {
		cur := l.tail
		for i := size - 1; i > index; i-- {
			cur = cur.prev
		}
		return cur.data, true
	}
}

func (l *DoubleLinkedList[T]) Length() int {
	return l.length
}

//
func (l *DoubleLinkedList[T]) Values() []T {
	values := []T{}

	for cur := l.head; cur != nil; cur = cur.next {
		values = append(values, cur.data)
	}
	return values
}
