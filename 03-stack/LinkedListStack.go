package _3_stack

import "fmt"

type node struct {
	data interface{}
	next *node
}

type linkedListStack struct {
	top *node
}

func NewLinkListStack() *linkedListStack {
	return  &linkedListStack{nil}
}

func NewNode(v interface{}) *node {
	return &node{
		data: v,
		next: nil,
	}
}

func (stack *linkedListStack) Push(v interface{}) {
	newNode := NewNode(v)
	if !stack.IsEmpty() {
		newNode.next = stack.top
	}
	stack.top = newNode
}

func (stack *linkedListStack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	top := stack.top.data
	stack.top = stack.top.next
	return top
}

func (stack *linkedListStack) IsEmpty() bool {
	return stack.top == nil
}

func (stack *linkedListStack) Top() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	return stack.top.data
}

func (stack *linkedListStack) Flush() {
	stack.top = nil
}

func (stack *linkedListStack) Print() {
	if stack.IsEmpty() {
		fmt.Println("empty stack")
	}
	cur := stack.top
	for cur != nil {
		fmt.Println(cur.data)
		cur = cur.next
	}
}