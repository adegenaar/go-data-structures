package containers

import (
	"reflect"
	"sort"
	"testing"
)

func TestBinaryTree_InsertAndTraverse(t *testing.T) {
	expected := []string{"d", "b", "g", "g", "c", "e", "a", "h", "f", "i", "j", "l", "k"}

	tree := &Tree[string]{}
	tree.PrettyPrint()
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

func TestBinaryTree_Walk(t *testing.T) {
	expected := []string{"d", "b", "g", "g", "c", "e", "a", "h", "f", "i", "j", "l", "k"}

	tree := &Tree[string]{}

	// check for walking the empty tree
	values := []string{}
	tree.Walk(func(n string) { values = append(values, n) })
	if len(values) != 0 {
		t.Errorf("Expected %v, got %v", 0, len(values))
	}

	// insert the expected values
	for i := 0; i < len(expected); i++ {
		tree.Insert(expected[i])
	}

	// walk the tree and validate the results
	values = []string{}
	tree.Walk(func(n string) { values = append(values, n) })
	sort.Strings(expected)
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}
func TestBinaryTree_Search(t *testing.T) {
	expected := []string{"d", "b", "g", "g", "c", "e", "a", "h", "f", "i", "j", "l", "k"}

	tree := &Tree[string]{}
	v := tree.Search("x")
	if v != nil {
		t.Errorf("Expected %v, got %v", nil, v)
	}
	for i := 0; i < len(expected); i++ {
		tree.Insert(expected[i])
	}
	values := []string{}
	for i := 0; i < len(expected); i++ {
		values = append(values, *tree.Search(expected[i]))
	}
	sort.Strings(values)
	sorted := make([]string, len(expected))
	copy(sorted, expected)

	sort.Strings(sorted)
	if !reflect.DeepEqual(values, sorted) {
		t.Errorf("Expected %v, got %v", sorted, values)
	}

	values = []string{}
	for i := len(expected) - 1; i >= 0; i-- {
		values = append(values, *tree.Search(expected[i]))
	}
	sort.Strings(values)
	sort.Strings(expected)
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

func TestBinaryTree_Delete(t *testing.T) {
	expected := []string{"d", "b", "g", "g", "c", "e", "a", "h", "f", "i", "j", "l", "k"}

	tree := &Tree[string]{}
	ret := tree.Delete("x")
	if ret == nil {
		t.Errorf("Expected %v, got %v", "delete not allowed on an empty tree", ret)
	}
	tree.Insert(expected[0])
	ret = tree.Delete("x")
	if ret == nil {
		t.Errorf("Expected %v, got %v", "value does not exist", ret)
	}
	ret = tree.Delete(expected[0])
	if ret != nil {
		t.Errorf("Expected %v, got %v", nil, ret)
	}

	tree.Insert(expected[0])
	tree.Insert(expected[1])
	ret = tree.Delete(expected[0])
	if ret != nil {
		t.Errorf("Expected %v, got %v", nil, ret)
	}
	ret = tree.Delete(expected[1])
	if ret != nil {
		t.Errorf("Expected %v, got %v", nil, ret)
	}
	tree.Insert(expected[0])
	tree.Insert(expected[1])
	ret = tree.Delete(expected[1])
	if ret != nil {
		t.Errorf("Expected %v, got %v", nil, ret)
	}
	ret = tree.Delete(expected[0])
	if ret != nil {
		t.Errorf("Expected %v, got %v", nil, ret)
	}
	for i := 0; i < len(expected); i++ {
		tree.Insert(expected[i])
	}
	for i := 0; i < len(expected); i++ {
		if err := tree.Delete(expected[i]); err != nil {
			t.Errorf("Expected %v, got %v", nil, err)
		}
	}
	if tree.length != 0 {
		t.Errorf("Expected %v, got %v", 0, tree.length)
	}
}

func TestBinaryTree_NilChecks(t *testing.T) {
	var tree *Tree[string] = nil

	tree.Insert("x")
	err := tree.Delete("x")
	if err == nil {
		t.Errorf("Expected %v, got %v", "delete not allowed on an empty tree", err)
	}
	tree.Search("x")
	tree.Walk(func(n string) {})
	tree.PrettyPrint()

	var node *binaryNode[string] = nil

	ret, _ := node.findMax(node)
	if ret != nil {
		t.Errorf("Expected %v, got %v", nil, ret)
	}
	err = node.replaceNode(node, nil)
	if err == nil {
		t.Errorf("Expected %v, got %v", "replaceNode() not allowed on a nil node", err)
	}

	node = node.searchNode("x")
	if node != nil {
		t.Errorf("Expected %v, got %v", nil, node)
	}

	node.addNode(node)
}
