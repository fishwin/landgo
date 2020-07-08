package shangguigu

import (
	"errors"
	"fmt"
)

type Stack struct {
	top int
	core []int
	capacity int
}

func NewStack(capacity int) *Stack {
	s := &Stack{
		top: -1,
		capacity: capacity,
		core: make([]int, capacity),
	}

	return s
}

func (s *Stack) Push(v int) error {
	if s.IsFull() {
		return errors.New("stack overflow")
	}
	s.top++
	s.core[s.top] = v
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("stack empty")
	}

	ret := s.core[s.top]
	s.top--
	return ret, nil
}

func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("stack empty")
	}
	return s.core[s.top], nil
}

func (s *Stack) IsFull() bool {
	return s.top == s.capacity - 1
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) Print() {
	for i := 0; i <= s.top ;i++ {
		fmt.Printf("%v\t", s.core[i])
	}
	fmt.Println()
}


