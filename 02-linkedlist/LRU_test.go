package main

import (
	"testing"
)
/**
单链表实现LRU单元测试
*/
func TestLRU_Read(t *testing.T) {
	lru := NewLRU(5)
	lru.Read(1)
	lru.Print()
	lru.Read(2)
	lru.Print()
	lru.Read(3)
	lru.Print()
	lru.Read(4)
	lru.Print()
	lru.Read(5)
	lru.Print()
	lru.Read(6)
	lru.Print()
	lru.Read(2)
	lru.Print()
	lru.Read(4)
	lru.Print()
}
