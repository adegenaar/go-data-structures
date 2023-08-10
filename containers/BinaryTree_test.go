package containers

import (
	"reflect"
	"sort"
	"testing"
)

func TestBinaryTree_InsertAndTraverse(t *testing.T) {
	expected := []string{"d", "b", "g", "g", "c", "e", "a", "h", "f", "i", "j", "l", "k"}

	tree := &Tree[string]{}
	for i := 0; i < len(expected); i++ {
		tree.Insert(expected[i])
	}
	tree.PrettyPrint()

	values := []string{}
	tree.Walk(func(n string) { values = append(values, n) })
	sort.Strings(expected)
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}
