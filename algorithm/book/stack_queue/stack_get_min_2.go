package stack_queue

import "github.com/fishwin/landgo/algorithm/book/base"

type GetMinStack2 struct {
	stackNormal *base.Stack
	stackMin    *base.Stack
}

func NewGetMinStack2() GetMinStack2 {
	return GetMinStack2{
		stackNormal: base.NewStack(),
		stackMin:    base.NewStack(),
	}
}

func (ms GetMinStack2) Push(v int) {
	ms.stackNormal.Push(v)

	if ms.stackMin.IsEmpty() {
		ms.stackMin.Push(v)
	} else {
		top, _ := ms.stackMin.Top()
		if v >= top {
			ms.stackMin.Push(top)
		} else {
			ms.stackMin.Push(v)
		}
	}
}

func (ms GetMinStack2) Pop() (int, error) {
	ret, err := ms.stackNormal.Pop()
	if err != nil {
		return ret, err
	}

	_, _ = ms.stackMin.Pop()
	return ret, nil
}

func (ms GetMinStack2) GetMin() (int, error) {
	return ms.stackMin.Top()
}
