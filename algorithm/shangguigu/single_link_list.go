package shangguigu

import (
	"errors"
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

type SingleLinkList struct {
	head   *Node
	tail   *Node
	length int
}

func NewSingleLinkList() *SingleLinkList {
	s := &SingleLinkList{
		head:   nil,
		length: 0,
	}
	return s
}

func (s *SingleLinkList) AddToTail(v int) {
	n := &Node{
		Val:  v,
		Next: nil,
	}

	if s.IsEmpty() {
		s.head = n
	} else {
		s.tail.Next = n
	}
	s.tail = n

	s.length++
}

func (s *SingleLinkList) AddToHead(v int) {
	n := &Node{
		Val:  v,
		Next: s.head,
	}

	if s.IsEmpty() {
		s.tail = n
	}
	s.head = n

	s.length++
}

func (s *SingleLinkList) AddAtIndex(index int, v int) error {
	if index < 0 || index > s.length {
		return errors.New("index overflow")
	}

	if index == 0 {
		s.AddToHead(v)
		return nil
	} else if index == s.length {
		s.AddToTail(v)
		return nil
	}
	// 需找到前一个节点的指针
	p := s.head
	for p != nil {
		index--
		if index == 0 {
			break
		}
		p = p.Next
	}

	n := &Node{
		Val: v,
	}

	n.Next = p.Next
	p.Next = n

	s.length++
	return nil
}

func (s *SingleLinkList) DeleteTail() error {
	if s.IsEmpty() {
		return errors.New("link list is empty")
	}

	if s.length == 1 {
		s.head = nil
		s.tail = nil
	} else {
		p := s.head
		for p.Next != s.tail {
			p = p.Next
		}

		p.Next = nil
		s.tail = p
	}
	s.length--
	return nil
}

func (s *SingleLinkList) DeleteHead() error {
	if s.IsEmpty() {
		return errors.New("link list is empty")
	}

	if s.length == 1 {
		s.head = nil
		s.tail = nil
	} else {
		s.head = s.head.Next
	}

	s.length--
	return nil
}

func (s *SingleLinkList) DeleteAtIndex(index int) error {
	if index < 0 || index > s.length-1 {
		return errors.New("index overflow")
	}
	if s.IsEmpty() {
		return errors.New("link list is empty")
	}

	if index == 0 {
		s.DeleteHead()
		return nil
	} else if index == s.length-1 {
		s.DeleteTail()
		return nil
	}

	// 需找到前一个节点的指针
	p := s.head
	for p != nil {
		index--
		if index == 0 {
			break
		}
		p = p.Next
	}

	temp := p.Next.Next
	p.Next.Next = nil
	p.Next = temp

	s.length--
	return nil
}

func (s *SingleLinkList) UpdateTail(v int) error {
	if s.IsEmpty() {
		return errors.New("link list is empty")
	}

	s.tail.Val = v
	return nil
}

func (s *SingleLinkList) UpdateHead(v int) error {
	if s.IsEmpty() {
		return errors.New("link list is empty")
	}

	s.head.Val = v
	return nil
}

func (s *SingleLinkList) UpdateAtIndex(index int, v int) error {
	if index < 0 || index > s.length-1 {
		return errors.New("index overflow")
	}
	if s.IsEmpty() {
		return errors.New("link list is empty")
	}

	p := s.head

	for p != nil {
		if index == 0 {
			break
		}
		index--
		p = p.Next
	}
	p.Val = v

	return nil
}

func (s *SingleLinkList) IsEmpty() bool {
	return s.head == s.tail && s.tail == nil
}

func (s *SingleLinkList) Length() int {
	return s.length
}

func (s *SingleLinkList) Print() {
	p := s.head
	for p != nil {
		fmt.Printf("%v ", p.Val)
		p = p.Next
	}
	fmt.Println()
	return
}
