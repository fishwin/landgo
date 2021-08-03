package stack_queue

import "github.com/fishwin/landgo/algorithm/book/base"

type GetMinStack struct {
	stackNormal *base.Stack
	stackMin    *base.Stack
}

func NewGetMinStack() GetMinStack {
	return GetMinStack{
		stackNormal: base.NewStack(),
		stackMin:    base.NewStack(),
	}
}

func (ms GetMinStack) Pop() (int, error) {

	ret, err := ms.stackNormal.Pop()
	if err != nil {
		return ret, err
	}

	if !ms.stackMin.IsEmpty() {
		top, _ := ms.stackMin.Top()
		if top == ret {
			_, _ = ms.stackMin.Pop()
		}
	}

	return ret, nil
}

func (ms GetMinStack) Push(v int) {
	ms.stackNormal.Push(v)

	if ms.stackMin.IsEmpty() {
		ms.stackMin.Push(v)
	} else {
		top, _ := ms.stackMin.Top()
		if v <= top {
			ms.stackMin.Push(v)
		}
	}
}

func (ms GetMinStack) GetMin() (int, error) {
	return ms.stackMin.Top()
}
