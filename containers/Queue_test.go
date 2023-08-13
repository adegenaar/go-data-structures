package containers

import (
	"reflect"
	"testing"
)

func TestQueue_EnqueueAndDequeue(t *testing.T) {
	s := NewQueue[int]()

	expected := []int{4, 3, 2, 1}
	for i := 0; i < len(expected); i++ {
		s.Enqueue(expected[i])
	}

	values := []int{}
	for {
		v, ok := s.Dequeue()
		if !ok {
			break
		}

		values = append(values, v)
	}

	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

func TestQueue_CheckNil(t *testing.T) {
	var q *Queue[int] = nil
	q.Enqueue(1)
	q.Dequeue()

	q = NewQueue[int]()

	expected := []int{4, 3, 2, 1}
	for i := 0; i < len(expected); i++ {
		q.Enqueue(expected[i])
	}

	values := []int{}
	for {
		v, ok := q.Dequeue()
		if !ok {
			break
		}

		values = append(values, v)
	}

	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}
