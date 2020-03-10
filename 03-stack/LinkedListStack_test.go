package _3_stack

import "testing"

func TestLinkedListStack_Push(t *testing.T) {
	stack := NewLinkListStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Print()
}

func TestLinkedListStack_Pop(t *testing.T) {
	stack := NewLinkListStack()
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