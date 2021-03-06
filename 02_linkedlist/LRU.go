package main

import "fmt"

/**
LRU 缓存淘汰机制单链表实现（非最佳实现）
LRU:LRU 是 Least Recently Used 的缩写，这种算法认为最近使用的数据是热门数据，下一次很大概率将会再次被使用。
    而最近很少被使用的数据，很大概率下一次不再用到。当缓存容量的满时候，优先淘汰最近很少使用的数据。
*/

type LRUNode struct {
	data interface{}
	next *LRUNode
}

type LRU struct {
	head      *LRUNode
	length    int
	maxLength int
}

func NewLRUNode(v interface{}) *LRUNode {
	return &LRUNode{
		data: v,
		next: nil,
	}
}

func NewLRU(max int) *LRU {
	return &LRU{
		head:      NewLRUNode(0),
		length:    0,
		maxLength: max,
	}
}

//查找缓存
func (list *LRU) Find(v interface{}) (pre, cur *LRUNode) {
	if list.length == 0 {
		return nil, nil
	}

	pre = list.head
	cur = list.head.next
	for cur != nil {
		if cur.data == v {
			break
		}
		pre = cur
		cur = cur.next
	}

	return
}

func (list *LRU) Delete(pre, cur *LRUNode) bool {
	if cur == nil {
		return false
	}

	pre.next = cur.next
	list.length--
	return true
}

func (list *LRU) InsertToHead(node *LRUNode) bool {
	//缓存空间已满
	if list.length == list.maxLength {
		return false
	}

	oldNode := list.head.next
	node.next = oldNode
	list.head.next = node
	list.length++
	return true
}

//删除最少用的缓存
func (list *LRU) DeleteTail() {
	pre := list.head
	cur := list.head.next
	for cur.next != nil {
		pre = cur
		cur = cur.next
	}
	pre.next = nil
	list.length--
}

func (list *LRU) Read(v interface{}) {
	pre, cur := list.Find(v)
	
	// 缓存存在
	if cur != nil {
		list.Delete(pre, cur)
		list.InsertToHead(cur)
		return
	}

	//缓存空间已满，删除最少访问的缓存
	if list.length >= list.maxLength {
		list.DeleteTail()
	}
	list.InsertToHead(NewLRUNode(v))
	return
}

func (list *LRU) Print() {
	cur := list.head
	str := ""
	for cur != nil {
		str += fmt.Sprintf("%+v", cur.data)
		cur = cur.next
		if cur != nil {
			str += "->"
		}
	}

	fmt.Println(str)
}
