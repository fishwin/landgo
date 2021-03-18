package shangguigu

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {

	stack := NewStack(5)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println(stack.Top())

	stack.Push(4)
	stack.Push(5)

	stack.Print()

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	stack.Push(6)
	stack.Push(7)
	fmt.Println(stack.Top())
	stack.Print()
	fmt.Println(stack.Pop())
	fmt.Println(stack.Top())
	stack.Print()

}
