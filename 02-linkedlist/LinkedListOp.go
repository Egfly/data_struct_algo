package main

import "fmt"

/**
带头单链表操作：反转、判断是否有环、合并两个有序单链表
*/

//反转
func (list *LinkedList) Reverse() {
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

/**
两个有序单链表合并
*/
func mergeSortedList(l1, l2 *LinkedList) *LinkedList {
	if l1 == nil || l1.Head == nil || l1.Head.Next == nil {
		return l2
	}
	if l2 == nil || l2.Head == nil || l2.Head.Next == nil {
		return l1
	}
	list := &LinkedList{Head: &LinkedNode{}}
	cur := list.Head
	cur1 := l1.Head.Next
	cur2 := l2.Head.Next
	for cur1 != nil && cur2 != nil {
		if cur1.Data.(int) > cur2.Data.(int) {
			cur.Next = cur2
			cur2 = cur2.Next
		} else {
			cur.Next = cur1
			cur1 = cur1.Next
		}
		cur = cur.Next
	}
	if cur1 != nil {
		cur.Next = cur1
	}
	if cur2 != nil {
		cur.Next = cur2
	}
	list.Length = l1.Length + l2.Length
	return list
}

/**
删除倒数第N个节点
*/
func (list *LinkedList) DeleteBottomN(n int) bool {
	if list.Head == nil || list.Head.Next == nil {
		return false
	}
	if n <= 0 || n > list.Length {
		return false
	}
	fast := list.Head.Next
	for i := 0; i < n && fast != nil; i++ {
		fast = fast.Next
	}

	if fast == nil {
		return false
	}
	slow := list.Head.Next
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return true
}

/**
查找中间节点
*/
func (list *LinkedList) FindMiddleNode() *LinkedNode {
	if list.Head == nil || list.Head.Next == nil {
		return  nil
	}
	if list.Head.Next.Next == nil {
		return list.Head.Next
	}
	slow, fast := list.Head, list.Head
	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func main() {
	list := NewLinkedList()
	list.InsertToTail(1)
	list.InsertToTail(3)
	list.InsertToTail(5)
	list.InsertToTail(7)
	list.InsertToTail(9)
	list.Print()
	//fmt.Println(list.HasCycle())
	//list.Head.Next.Next.Next.Next.Next.Next = list.Head.Next.Next
	//fmt.Println(list.HasCycle())

	list1 := NewLinkedList()
	list1.InsertToTail(2)
	list1.InsertToTail(4)
	list1.InsertToTail(6)
	list1.InsertToTail(8)
	list1.InsertToTail(10)
	list1.InsertToTail(12)
	list1.Print()
	//
	//res := mergeSortedList(list1, list)
	//res.Print()
	//res.DeleteBottomN(2)
	//res.Print()
	//fmt.Println(res.FindMiddleNode())

	fmt.Println(list.FindMiddleNode())
	fmt.Println(list1.FindMiddleNode())
}
