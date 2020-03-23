package _3_stack

import "fmt"

type arrayStack struct {
	data  []interface{}
	top   int
}

func NewArrayStack(max int) *arrayStack {
	return &arrayStack{
		data: make([]interface{}, max),
		top:  -1,
	}
}

func (stack *arrayStack) Push(v interface{}) {
	stack.top++
	if stack.top > len(stack.data) - 1 {
		stack.data = append(stack.data, v)
	} else {
		stack.data[stack.top] = v
	}
}

func (stack *arrayStack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	v := stack.data[stack.top]
	stack.data[stack.top] = nil
	stack.top -= 1
	return  v
}

func (stack *arrayStack) IsEmpty() bool {
	if stack.top < 0 {
		return  true
	}
	return false
}

func (stack *arrayStack) Top() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	return stack.data[stack.top]
}

func (stack *arrayStack) Flush() {
	stack.top = -1;
	for i, _ := range stack.data {
		stack.data[i] = nil
	}
}

func (stack *arrayStack) Print() {
	if stack.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		for i := stack.top; i >= 0; i-- {
			fmt.Println(stack.data[i])
		}
	}
}
