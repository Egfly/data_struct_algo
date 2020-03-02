package main

import "testing"
/**
单链表实现LRU单元测试
*/
func TestLRU_Read(t *testing.T) {
	lru := NewLRU(10)
	lru.Read(1)
}
