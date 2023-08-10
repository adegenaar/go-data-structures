package containers

import (
	"reflect"
	"testing"
)

func TestDoubleLinked_InsertAtHead(t *testing.T) {
	link := NewDoubleLinkedList[int]()
	link.InsertAtHead(1)
	link.InsertAtHead(2)
	link.InsertAtHead(3)
	link.InsertAtHead(4)

	expected := []int{4, 3, 2, 1}
	values := link.Values()
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

// test with a string instead of an int
func TestDoubleLinked_InsertAtHeadString(t *testing.T) {
	link := NewDoubleLinkedList[string]()
	link.InsertAtHead("one")
	link.InsertAtHead("two")
	link.InsertAtHead("three")
	link.InsertAtHead("four")

	expected := []string{"four", "three", "two", "one"}
	values := link.Values()
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

func TestDoubleLinked_InsertAtTail(t *testing.T) {
	link := NewDoubleLinkedList[int]()
	link.InsertAtTail(1)
	link.InsertAtTail(2)
	link.InsertAtTail(3)
	link.InsertAtTail(4)

	expected := []int{1, 2, 3, 4}
	values := link.Values()
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

func TestDoubleLinked_InsertAt(t *testing.T) {
	link := NewDoubleLinkedList[int]()
	link.InsertAtTail(1)
	link.InsertAtTail(2)
	link.InsertAtTail(3)
	link.InsertAtTail(4)

	expected := []int{1, 2, 3, 4}
	values := link.Values()
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}

	// try and insert before the list
	err := link.InsertAt(-1, 0)
	if err == nil {
		t.Errorf("InsertAt(-1): Expected %v, got %v", nil, err)
	}

	// insert at the head of the list
	link.InsertAt(0, 0)
	value, _ := link.Get(0)
	if value != 0 {
		t.Errorf("Get(-1): Expected %v, got %v", 0, value)
	}
	expected = []int{0, 1, 2, 3, 4}
	values = link.Values()
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
	link.InsertAt(1, 11)
	value, _ = link.Get(1)
	if value != 11 {
		t.Errorf("Get(-1): Expected %v, got %v", 0, value)
	}
	expected = []int{0, 11, 1, 2, 3, 4}
	values = link.Values()
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
	link.InsertAt(3, 6)
	link.InsertAt(6, 9)
	link.InsertAt(link.Length(), 12)

	expected = []int{0, 11, 1, 6, 2, 3, 9, 4, 12}
	values = link.Values()
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

func TestDoubleLinked_DeleteAt(t *testing.T) {
	link := NewDoubleLinkedList[int]()

	// handle the special cases - deleting the head of an empty list
	err := link.DeleteAtHead()
	if err == nil {
		t.Errorf("DeleteAtHead: Expected %v, got %v", "list is empty", err)
	}
	err = link.DeleteAt(-1)
	if err == nil {
		t.Errorf("DeleteAt: Expected %v, got %v", "index out of range", err)
	}
	err = link.DeleteAt(5)
	if err == nil {
		t.Errorf("DeleteAt: Expected %v, got %v", "index out of range", err)
	}
	link.InsertAtTail(1) // { 1 }
	link.InsertAtTail(2) // { 1, 2 }
	link.InsertAtTail(3) // { 1, 2, 3 }
	link.InsertAtTail(4) // { 1, 2, 3, 4 }

	link.DeleteAt(0) // delete the 1 - { 2, 3 ,4}
	expected := []int{2, 3, 4}
	values := link.Values()
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}

	link.DeleteAt(2) // delete the 4 - { 2, 3 }
	expected = []int{2, 3}
	values = link.Values()
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}

	link.InsertAtHead(1) // { 1, 2, 3 }
	link.InsertAtTail(4) // { 1, 2, 3, 4 }
	link.InsertAtTail(5) // { 1, 2, 3, 4, 5 }
	link.InsertAtTail(6) // { 1, 2, 3, 4, 5, 6 }
	link.InsertAtTail(7) // { 1, 2, 3, 4, 5, 6, 7 }
	link.DeleteAt(2)     // delete the 3 { 1, 2, 4, 5, 6, 7 }
	expected = []int{1, 2, 4, 5, 6, 7}
	values = link.Values()
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
	link.DeleteAtTail() // { 1, 2, 4, 5, 6 }
	link.DeleteAtTail() // { 1, 2, 4, 5 }

	link.DeleteAt(3) // delete the 5 { 1, 2, 4 }
	link.DeleteAt(1) // delete the 2 { 1, 4 }
	expected = []int{1, 4}
	values = link.Values()
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}

	link.DeleteAtTail() // delete the 4 - {1}

	expected = []int{1}
	values = link.Values()
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}

}

func TestDoubleLinked_DeleteAtTail(t *testing.T) {
	link := NewDoubleLinkedList[int]()
	err := link.DeleteAtTail()
	if err == nil {
		t.Errorf("DeleteAtTail: Expected %v, got %v", "list is empty", err)
	}
	link.InsertAtTail(1)

	//special case: DeleteAtTail with a list of one
	err = link.DeleteAtTail()
	if err != nil {
		t.Errorf("DeleteAtTail: Expected %v, got %v", nil, err)
	}

	link.InsertAtTail(1)
	link.InsertAtTail(2)
	link.InsertAtTail(3)

	// general case:
	err = link.DeleteAtTail()
	if err != nil {
		t.Errorf("DeleteAtTail: Expected %v, got %v", nil, err)
	}

	expected := []int{1, 2}
	values := link.Values()
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

func TestDoubleLinked_Get(t *testing.T) {
	link := NewDoubleLinkedList[int]()
	link.InsertAtHead(1)
	link.InsertAtHead(2)
	link.InsertAtHead(3)
	link.InsertAtHead(4)

	expected := []int{4, 3, 2, 1}
	values := link.Values()
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
	_, failure := link.Get(-1)
	if failure {
		t.Errorf("Get(-1): Expected %v, got %v", false, failure)
	}

	value, success := link.Get(2)
	if !success {
		t.Errorf("Expected %v, got %v", true, success)
	}
	if value != 2 {
		t.Errorf("Expected %v, got %v", 2, value)
	}
}
