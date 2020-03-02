package main

import "fmt"

type LinkedNode struct {
	Data interface{}
	Next *LinkedNode
}

type LinkedList struct {
	Head   *LinkedNode
	Length int
}

func NewLinkedNode(v interface{}) *LinkedNode {
	return &LinkedNode{
		Data: v,
		Next: nil,
	}
}

func (node *LinkedNode) GetData() interface{} {
	return node.Data
}

func (node *LinkedNode) GetNext() *LinkedNode {
	return node.Next
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head:   NewLinkedNode(0),
		Length: 0,
	}
}

//某节点之后插入节点
func (list *LinkedList) InsertAfter(p *LinkedNode, v interface{}) bool {
	if p == nil {
		return false
	}
	newNode := NewLinkedNode(v)
	oldNext := p.Next
	p.Next = newNode
	newNode.Next = oldNext
	list.Length++
	return true
}

//某节点之前插入节点
func (list *LinkedList) InsertBefore(p *LinkedNode, v interface{}) bool {
	if p == nil || list.Head == p {
		return false
	}
	cur := list.Head.Next
	pre := list.Head
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.Next
	}
	if cur == nil {
		return false
	}
	newNode := NewLinkedNode(v)
	pre.Next = newNode
	newNode.Next = cur
	list.Length++
	return true
}

//在链表头部插入
func (list *LinkedList) InsertToHead(v interface{}) bool {
	return list.InsertAfter(list.Head, v)
}

//在链表尾部插入
func (list *LinkedList) InsertToTail(v interface{}) bool {
	cur := list.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	return list.InsertAfter(cur, v)
}

//通过索引查找节点
func (list *LinkedList) FindByIndex(index int) *LinkedNode {
	if index >= list.Length {
		return nil
	}
	p := 0
	cur := list.Head
	for p < list.Length {
		if p == index {
			break
		}
		p++
		cur = cur.Next
	}
	return  cur
}

//删除节点
func (list *LinkedList) DeleteNode(p *LinkedNode) bool {
	if p == nil {
		return false
	}
	cur := list.Head.Next
	pre := list.Head
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.Next
	}
	if cur == nil {
		return false
	}
	pre.Next = cur.Next
	list.Length--
	return true
}

//打印链表
func (list *LinkedList) Print() {
	cur := list.Head.Next
	format := ""
	for cur != nil {
		format +=fmt.Sprintf("%+v", cur.GetData())
		cur = cur.Next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}
