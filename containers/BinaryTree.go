package containers

import (
	"errors"
	"fmt"
	"strings"
)

// TODO: Figure out the "Ordered"
type binaryNode[T ~int | ~float64 | ~string] struct {
	data  T
	right *binaryNode[T]
	left  *binaryNode[T]
}

func newNode[T ~int | ~float64 | ~string](data T) *binaryNode[T] {
	n := new(binaryNode[T])
	n.data = data
	return n
}

func (n *binaryNode[T]) addNode(cur *binaryNode[T]) {
	if n == nil {
		return
	}
	if n.data < cur.data {
		if n.right == nil {
			n.right = cur
			return
		}
		n.right.addNode(cur)
	} else {
		if n.left == nil {
			n.left = cur
			return
		}
		n.left.addNode(cur)
	}
}

func (n *binaryNode[T]) searchNode(val T) *binaryNode[T] {
	if n == nil {
		return nil
	}

	if n.data < val {
		if n.right == nil {
			return nil
		}
		return n.right.searchNode(val)
	}
	if n.data > val {
		if n.left == nil {
			return nil
		}
		return n.left.searchNode(val)
	}

	// found it
	return n
}

func (t *binaryNode[T]) walk(n *binaryNode[T], f func(T)) {
	if n == nil {
		return
	}
	t.walk(n.left, f)
	f(n.data)
	t.walk(n.right, f)
}

func (n *binaryNode[T]) removeNode(value T, parent *binaryNode[T]) error {
	//Delete node from binary search tree
	if n == nil {
		return errors.New("value does not exist")
	}
	if value < n.data {
		return n.left.removeNode(value, n)
	}
	if value > n.data {
		return n.right.removeNode(value, n)
	}
	// must be equal...

	// check for the simple case where there are no children
	if n.left == nil && n.right == nil {
		n.replaceNode(parent, nil)
		return nil
	}

	// check for the case where there is only one child
	if n.left == nil {
		n.replaceNode(parent, n.right)
		return nil
	}
	if n.right == nil {
		n.replaceNode(parent, n.left)
		return nil
	}

	// there are two children, so replace by the maximum valued child
	replacement, replParent := n.left.findMax(n)
	n.data = replacement.data

	return replacement.removeNode(replacement.data, replParent)
}

// find the child node with the maximum value
func (n *binaryNode[T]) findMax(parent *binaryNode[T]) (*binaryNode[T], *binaryNode[T]) {
	if n == nil {
		return nil, parent
	}
	if n.right == nil {
		return n, parent
	}
	return n.right.findMax(n)
}

func (n *binaryNode[T]) replaceNode(parent, replacement *binaryNode[T]) error {
	if n == nil {
		return errors.New("replaceNode() not allowed on a nil node")
	}

	if n == parent.left {
		parent.left = replacement
		return nil
	}
	parent.right = replacement
	return nil
}

////////////////////////////////////////////////////////////////////////

type Tree[T ~int | ~float64 | ~string] struct {
	root   *binaryNode[T]
	length int
}

func (t *Tree[T]) Insert(value T) {
	if t == nil {
		return
	}
	n := newNode(value)
	if t.root == nil {
		t.root = n
	} else {
		t.root.addNode(n)
	}
	t.length++
}

func (t *Tree[T]) Delete(value T) error {
	if t == nil {
		return errors.New("delete not allowed on an uninitialized tree")
	}
	if t.root == nil {
		return errors.New("delete not allowed on an empty tree")
	}

	fakeParent := &binaryNode[T]{right: t.root}
	err := t.root.removeNode(value, fakeParent)
	if err != nil {
		return err
	}

	if fakeParent.right == nil {
		t.root = nil
	}
	t.length--
	return nil
}

func (t *Tree[T]) Search(value T) *T {
	if t == nil || t.root == nil {
		return nil
	}

	n := t.root.searchNode(value)
	return &n.data
}

func (t *Tree[T]) Walk(f func(T)) {
	if t == nil || t.root == nil {
		return
	}
	t.root.walk(t.root, f)
}

func (t *Tree[T]) PrettyPrint() {
	if t == nil || t.root == nil {
		return
	}
	printNode := func(n *binaryNode[T], depth int) {
		fmt.Printf("%s%v\n", strings.Repeat("  ", depth), n.data)
	}
	var walk func(*binaryNode[T], int)
	walk = func(n *binaryNode[T], depth int) {
		if n == nil {
			return
		}
		walk(n.right, depth+1)
		printNode(n, depth)
		walk(n.left, depth+1)
	}

	walk(t.root, 0)
}
