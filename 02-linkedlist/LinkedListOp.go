package main

import "fmt"

/**
带头单链表操作：反转、判断是否有环、合并两个有序单链表
*/

//反转
func (list *LinkedList) Reverse()  {
	if list.Head == nil || list.Head.Next == nil || list.Head.Next.Next == nil {
		return
	}
	var pre *LinkedNode = nil
	cur := list.Head.Next
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	list.Head.Next = pre
}

//单链表是否有环 (快、慢指针)
func (list *LinkedList) HasCycle() bool {
	if list.Head == nil || list.Head.Next == nil || list.Head.Next.Next == nil {
		return false
	}
	slow := list.Head
	fast := list.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func main()  {
	list := NewLinkedList()
	list.InsertAfter(list.Head, 1)
	list.InsertAfter(list.Head, 2)
	list.InsertAfter(list.Head, 3)
	list.InsertAfter(list.Head, 4)
	list.InsertAfter(list.Head, 5)
	list.Print()
	fmt.Println(list.HasCycle())
	list.Head.Next.Next.Next.Next.Next.Next = list.Head.Next.Next
	fmt.Println(list.HasCycle())
}