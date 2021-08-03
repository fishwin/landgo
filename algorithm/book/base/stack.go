package base

import (
	"errors"
)

var (
	ErrStackEmpty = errors.New("stack is empty")
)

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{
		data: make([]int, 0),
	}
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, ErrStackEmpty
	}

	index := len(s.data) - 1
	ret := s.data[index]
	s.data = s.data[:index]
	return ret, nil
}

func (s *Stack) Push(v int) {
	s.data = append(s.data, v)
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return -1, ErrStackEmpty
	}

	return s.data[len(s.data) - 1], nil
}

func (s *Stack) All() []int {
	return s.data
}


