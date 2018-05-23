package queue_test

import (
	"testing"

	"github.com/shivammg/ppds/queue"
)

type MyQueue struct {
	elems []int
}

func (q *MyQueue) push(ele int) {
	q.elems = append([]int{ele}, q.elems...)
}

func (q *MyQueue) pop() (int, bool) {
	if len(q.elems) == 0 {
		return 0, false
	}
	e := q.elems[len(q.elems)-1]
	q.elems = q.elems[:len(q.elems)-1]
	return e, true
}

func (q *MyQueue) Enqueue(ele interface{}) {
	q.push(ele.(int))
}

func (q *MyQueue) Dequeue() (interface{}, bool) {
	return q.pop()
}

func TestSprint(t *testing.T) {
	q := MyQueue{}
	q.push(10)
	q.push(11)
	q.push(12)
	q.push(13)
	want := ` ┌──┐ ┌──┐ ┌──┐ ┌──┐
→│10│→│11│→│12│→│13│→
 ├──┤ ├──┤ ├──┤ ├──┤
 │ 1│ │ 2│ │ 3│ │ 4│
 └──┘ └──┘ └──┘ └──┘
`
	got := queue.Sprint(&q)
	if got != want {
		t.Errorf("Expected: %s\nGot: %s", want, got)
	}
}
