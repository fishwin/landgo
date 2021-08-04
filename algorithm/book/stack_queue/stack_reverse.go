package stack_queue

import "github.com/fishwin/landgo/algorithm/book/base"

func GetAndRemoveBottomElement(stack *base.Stack) int {
	if stack == nil || stack.IsEmpty() {
		return -1
	}

	top, _ := stack.Pop()
	if stack.IsEmpty() {
		return top
	}

	last := GetAndRemoveBottomElement(stack)
	stack.Push(top)
	return last
}

func ReverseStack(stack *base.Stack) {
	if stack == nil || stack.IsEmpty() {
		return
	}

	bottom := GetAndRemoveBottomElement(stack)
	ReverseStack(stack)
	stack.Push(bottom)
}
