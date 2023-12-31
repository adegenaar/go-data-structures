package containers

import (
	"reflect"
	"testing"
)

func TestStack_pushandpop(t *testing.T) {

	s := NewStack[int]()

	expected := []int{4, 3, 2, 1}
	for i := 0; i < len(expected); i++ {
		s.Push(expected[i])
	}

	values := []int{}
	for {
		v, ok := s.Pop()
		if !ok {
			break
		}

		values = append(values, v)
	}
	reverseSlice(values)
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

func TestStack_CheckNil(t *testing.T) {
	var s *Stack[int] = nil
	s.Push(1)
	s.Pop()

	s = NewStack[int]()

	expected := []int{4, 3, 2, 1}
	for i := 0; i < len(expected); i++ {
		s.Push(expected[i])
	}

	values := []int{}
	for {
		v, ok := s.Pop()
		if !ok {
			break
		}

		values = append(values, v)
	}

	reverseSlice(values)

	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

// panic if s is not a slice
func reverseSlice(s interface{}) {
	size := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
