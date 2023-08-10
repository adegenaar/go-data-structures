package containers

import (
	"reflect"
	"testing"
)

func TestSingleLinked_InsertAtHead(t *testing.T) {
	link := NewSingleLinkedList[int]()
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
func TestSingleLinked_InsertAtHeadString(t *testing.T) {
	link := NewSingleLinkedList[string]()
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

func TestSingleLinked_InsertAtTail(t *testing.T) {
	link := NewSingleLinkedList[int]()
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

func TestSingleLinked_InsertAt(t *testing.T) {
	link := NewSingleLinkedList[int]()
	link.InsertAtTail(1)
	link.InsertAtTail(2)
	link.InsertAtTail(3)
	link.InsertAtTail(4)

	err := link.InsertAt(-1, 0)
	if err == nil {
		t.Errorf("InsertAt(-1): Expected %v, got %v", nil, err)
	}

	link.InsertAt(0, 0)
	link.InsertAt(1, 11)
	link.InsertAt(3, 6)
	link.InsertAt(link.Length(), 12)

	expected := []int{0, 11, 1, 6, 2, 3, 4, 12}
	values := link.Values()
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

func TestSingleLinked_DeleteAt(t *testing.T) {
	link := NewSingleLinkedList[int]()
	err := link.DeleteAtHead()
	if err == nil {
		t.Errorf("DeleteAtHead: Expected %v, got %v", "list is empty", err)
	}
	err = link.DeleteAt(-1)
	if err == nil {
		t.Errorf("DeleteAt: Expected %v, got %v", "index out of range", err)
	}
	link.InsertAtTail(1)
	link.InsertAtTail(2)
	link.InsertAtTail(3)
	link.InsertAtTail(4)

	link.DeleteAtHead()
	link.DeleteAt(2)
	link.DeleteAt(0)

	expected := []int{3}
	values := link.Values()
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}

}

func TestSingleLinked_DeleteAtTail(t *testing.T) {
	link := NewSingleLinkedList[int]()
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

func TestSingleLinked_Get(t *testing.T) {
	link := NewSingleLinkedList[int]()
	link.InsertAtHead(1)
	link.InsertAtHead(2)
	link.InsertAtHead(3)
	link.InsertAtHead(4)

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
