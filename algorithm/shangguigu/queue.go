package shangguigu

import (
	"errors"
	"fmt"
)

type Queue struct {
	core []int
	head int // 指向第一个元素
	tail int // 指向最后一个元素的后一个位置
}

func NewQueue(length int) *Queue {
	q := &Queue{
		core: make([]int, length+1),
		head: 0,
		tail: 0,
	}
	fmt.Println("test111")

	return q
}

func (q *Queue) EnQueue(v int) error {
	if q.IsFull() {
		return errors.New("queue is full")
	}

	q.core[q.tail] = v
	q.tail = (q.tail + 1) % len(q.core)
	return nil
}

func (q *Queue) DeQueue() (int, error) {
	if q.IsEmpty() {
		return -1, errors.New("queue is empty")
	}

	ret := q.core[q.head]
	q.head = (q.head + 1) % len(q.core)
	return ret, nil
}

func (q *Queue) Head() (int, error) {
	if q.IsEmpty() {
		return -1, errors.New("queue is empty")
	}

	return q.core[q.head], nil
}

func (q *Queue) Tail() (int, error) {
	if q.IsEmpty() {
		return -1, errors.New("queue is empty")
	}

	return q.core[(q.tail-1)%len(q.core)], nil
}

func (q *Queue) IsFull() bool {
	return (q.tail+1)%len(q.core) == q.head
}

func (q *Queue) IsEmpty() bool {
	return q.head == q.tail
}

func (q *Queue) Print() {
	for i := q.head; i%len(q.core) != q.tail; i++ {
		fmt.Printf("%v ", q.core[i%len(q.core)])
	}
	fmt.Println()
}
