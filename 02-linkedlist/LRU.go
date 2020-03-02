package main

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
	head   *LRUNode
	length int
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
		head:   NewLRUNode(nil),
		length: 0,
		maxLength: max,
	}
}

//查找缓存
func (list *LRU) Find(v interface{}) (pre, cur *LRUNode) {
	if list.length == 0 {
		return nil, nil
	}

	pre = nil
	cur = list.head
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
	if pre == nil {
		list.head = nil
		list.length = 0
	}

	pre.next = cur.next
	list.length--
	return true
}


func (list *LRU) InsertToHead(node *LRUNode) bool {
	//缓存空间为空
	if list.length == 0 {
		list.head = node
		list.length++
		return true
	}

	//缓存空间已满
	if list.length == list.maxLength {
		return false
	}
	newNode := node
	oldHead := list.head
	newNode.next = oldHead
	list.head = newNode
	list.length++
	return  true
}

//删除最少用的缓存
func (list *LRU) DeleteTail() {
	pre := list.head
	cur := list.head.next
	for cur != nil {
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
	}

	//缓存空间已满，删除最少访问的缓存
	if list.length >= list.maxLength {
		list.DeleteTail()
	}
	list.InsertToHead(NewLRUNode(v))
}