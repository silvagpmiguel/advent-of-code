package utils

import (
	"fmt"
	"strings"
)

// Node of a tree
type Node struct {
	Left  *Node
	Right *Node
	Val   interface{}
}

// Tree struct
type Tree struct {
	Node *Node
	Comparator
}

// Comparator func
type Comparator func(interface{}, interface{}) int

// NewTree creates a new tree ordered by a user defined comparator
func NewTree(comp Comparator, values ...interface{}) *Tree {
	t := &Tree{Comparator: comp}

	for _, val := range values {
		t.Node = t.Node.add(comp, val)
	}

	return t
}

// Add a value to the tree
func (t *Tree) Add(value interface{}) *Tree {
	t.Node.add(t.Comparator, value)
	return t
}

// Search a value of the tree
func (t *Tree) Search(value interface{}) *Node {
	return t.Node.search(t.Comparator, value)
}

// Remove TODO
func (t *Tree) Remove(value interface{}) *Tree {
	return t
}

// Reverse tree
func (t *Tree) Reverse() *Tree {
	t.Node.reverse()
	return t
}

// Clone makes a copy of a tree
func (t *Tree) Clone() *Tree {
	cloned := &Tree{Comparator: t.Comparator}
	t.Node.clone(t.Comparator, cloned)
	return cloned
}

// Length tree
func (t *Tree) Length(value interface{}) int {
	return t.Node.length()
}

// StringInOrder returns a string array in order
func (t *Tree) StringInOrder() []string {
	arr := []string{}
	t.Node.stringInOrder(&arr)
	return arr
}

// StringPreOrder returns a string array in pre order
func (t *Tree) StringPreOrder() []string {
	arr := []string{}
	t.Node.stringPreOrder(&arr)
	return arr
}

// StringPosOrder returns a string array in pos order
func (t *Tree) StringPosOrder() []string {
	arr := []string{}
	t.Node.stringPosOrder(&arr)
	return arr
}

// String returns the string method of this type
func (t *Tree) String() string {
	return strings.Join(t.StringInOrder(), ", ")
}

// String returns the string method of this type
func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Val)
}

func (n *Node) add(comp Comparator, val interface{}) *Node {
	if n == nil {
		return &Node{Val: val}
	}

	if comp(n.Val, val) > 0 {
		n.Left = n.Left.add(comp, val)
		return n
	}

	n.Right = n.Right.add(comp, val)
	return n
}

func (n *Node) reverse() {
	if n == nil {
		return
	}

	temp := n.Right
	n.Right = n.Left
	n.Left = temp

	n.Left.reverse()
	n.Right.reverse()
}

func (n *Node) clone(comp Comparator, t *Tree) {
	if n == nil {
		return
	}

	n.Left.clone(comp, t)
	t.Node = t.Node.add(comp, n.Val)
	n.Right.clone(comp, t)
}

func (n *Node) length() int {
	return 1 + n.Left.length() + n.Right.length()
}

func (n *Node) search(comp Comparator, val interface{}) *Node {
	if comp(n.Val, val) == 0 {
		return n
	} else if comp(n.Val, val) > 0 {
		return n.Left.search(comp, val)
	} else {
		return n.Right.search(comp, val)
	}
}

func (n *Node) stringInOrder(arr *[]string) {
	if n == nil {
		return
	}

	n.Left.stringInOrder(arr)
	*arr = append(*arr, fmt.Sprintf("%v", n.Val))
	n.Right.stringInOrder(arr)
}

func (n *Node) stringPreOrder(arr *[]string) {
	if n == nil {
		return
	}

	n.Left.stringPreOrder(arr)
	*arr = append(*arr, fmt.Sprintf("%v", n.Val))
	n.Right.stringPreOrder(arr)
}

func (n *Node) stringPosOrder(arr *[]string) {
	if n == nil {
		return
	}

	n.Left.stringPreOrder(arr)
	*arr = append(*arr, fmt.Sprintf("%v", n.Val))
	n.Right.stringPreOrder(arr)
}
