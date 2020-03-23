package _3_stack

import (
	"fmt"
	"testing"
)

func TestArrayStack_Push(t *testing.T) {
	stack := NewArrayStack(10)
	stack.Print()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Print()
}

func TestArrayStack_Pop(t *testing.T) {
	stack := NewArrayStack(10)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Print()
	stack.Pop()
	stack.Print()
	stack.Pop()
	stack.Print()
	stack.Pop()
	stack.Print()
}

func TestArrayStack_Top(t *testing.T) {
	stack := NewArrayStack(10)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	top := stack.Top()
	fmt.Println(top)
}

func TestArrayStack_Flush(t *testing.T) {
	stack := NewArrayStack(10)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Flush()
	stack.Print()
}