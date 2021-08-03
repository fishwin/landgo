package stack_queue

import "github.com/fishwin/landgo/algorithm/book/base"

type QueueVia2Stack struct {
	stackIn *base.Stack
	stackOut *base.Stack
}

func NewQueueVia2Stack() QueueVia2Stack {
	return QueueVia2Stack{
		stackIn:  base.NewStack(),
		stackOut: base.NewStack(),
	}
}

func (q QueueVia2Stack) EnQueue(v int) {
	//for !q.stackOut.IsEmpty() {
	//	t, _ := q.stackOut.Pop()
	//	q.stackIn.Push(t)
	//}
	//q.stackIn.Push(v)

	q.stackIn.Push(v)
}

func (q QueueVia2Stack) DeQueue() (int, error) {
	//for !q.stackIn.IsEmpty() {
	//	t, _ := q.stackIn.Pop()
	//	q.stackOut.Push(t)
	//}
	//return q.stackOut.Pop()

	if !q.stackOut.IsEmpty() {
		return q.stackOut.Pop()
	}

	for !q.stackIn.IsEmpty() {
		t, _ := q.stackIn.Pop()
		q.stackOut.Push(t)
	}
	return q.stackOut.Pop()
}
