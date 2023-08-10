package containers

import (
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

func (t *binaryNode[T]) walk(n *binaryNode[T], f func(T)) {
	if n == nil {
		return
	}
	t.walk(n.left, f)
	f(n.data)
	t.walk(n.right, f)
}

////////////////////////////////////////////////////////////////////////

type Tree[T ~int | ~float64 | ~string] struct {
	root   *binaryNode[T]
	length int
}

func (t *Tree[T]) Insert(value T) {
	n := newNode(value)
	if t.root == nil {
		t.root = n
	} else {
		t.root.addNode(n)
	}
	t.length++
}

func (t *Tree[T]) Walk(f func(T)) {
	t.root.walk(t.root, f)
}

func (t *Tree[T]) PrettyPrint() {
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
