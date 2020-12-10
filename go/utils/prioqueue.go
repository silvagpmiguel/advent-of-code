package utils

import (
	"fmt"
	"strings"
)

// PQueue structure
type PQueue struct {
	values []int
	length int
	isMin  bool
}

// NewPQueue creates a new priority queue
func NewPQueue(isMin bool) *PQueue {
	return &PQueue{
		values: []int{},
		length: 0,
		isMin:  isMin,
	}
}

// Push element to the queue
func (q *PQueue) Push(val ...int) *PQueue {
	for _, val := range val {
		q.values = append(q.values, val)
		q.bubbleUp(q.length)
		q.length++
	}

	return q
}

// Length of the queue
func (q *PQueue) Length() int {
	return q.length
}

// IsEmpty checks if the queue is empty
func (q *PQueue) IsEmpty() bool {
	return q.length == 0
}

// Peek element to queue
func (q *PQueue) Peek() (int, error) {
	if q.IsEmpty() {
		return -1, fmt.Errorf("stack is empty")
	}

	return q.values[0], nil
}

// Pop element from the queue
func (q *PQueue) Pop() (int, error) {
	if q.IsEmpty() {
		return -1, fmt.Errorf("stack is empty")
	}

	last := q.length - 1
	q.swap(0, last)
	toRemove := q.values[last]
	q.values = q.values[:last]
	q.length--
	q.bubbleDown(0, last)

	return toRemove, nil
}

// Clone queue
func (q *PQueue) Clone() *PQueue {
	return &PQueue{
		values: q.values,
		length: q.length,
		isMin:  q.isMin,
	}
}

// String method
func (q *PQueue) String() string {
	qType := "Max Priority Queue"

	if q.isMin {
		qType = "Min Priority Queue"
	}

	return fmt.Sprintf("%s (len %v)\n%v", qType, q.length, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(q.values)), ","), "[]"))
}

/** Aux Methods */

func (q *PQueue) bubbleDown(start int, end int) {
	length := q.length

	for i := start; ; {
		child := q.leftChild(i)

		if child >= length {
			break
		}

		if rightChild := child + 1; rightChild < length && q.needSwap(rightChild, child) {
			child = rightChild
		}

		if !q.needSwap(child, i) {
			break
		}

		q.swap(i, child)
		i = child
	}
}

func (q *PQueue) bubbleUp(actual int) {
	for i := actual; i >= 0; {
		parent := q.parent(i)

		if !q.needSwap(i, parent) {
			break
		}

		q.swap(i, parent)
		i = parent
	}
}

func (q *PQueue) parent(i int) int {
	return (i - 1) / 2
}

func (q *PQueue) rightChild(i int) int {
	return 2*i + 2
}

func (q *PQueue) leftChild(i int) int {
	return 2*i + 1
}

func (q *PQueue) needSwap(i int, j int) bool {
	if q.isMin {
		return q.values[i] < q.values[j]
	}

	return q.values[i] > q.values[j]
}

func (q *PQueue) swap(i int, j int) {
	aux := q.values[i]
	q.values[i] = q.values[j]
	q.values[j] = aux
}
