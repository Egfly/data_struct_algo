package _3_stack

type Stack interface {
	Push(item interface{})
	Pop() interface{}
	IsEmpty() bool
	Top() interface{}
	Flush()
	Print()
}
