package sort

import (
	"github.com/cheekybits/genny/generic"
	"sync"
)

type Item generic.Type

type Stack struct {
	// 栈中的数据
	items []Item
	// 锁
	lock sync.RWMutex
}

/**
 * 创建栈
 */
func (s *Stack)New(item Item) * Stack {
	s.items = []Item{}
	return s
}

/**
 * 入栈
 */
func (s *Stack)Push(item Item) {
	s.lock.Lock()
	s.items = append(s.items, item)
	s.lock.Unlock()
}

/**
 * 出栈
 */
func (s *Stack)Pop() * Item {
	s.lock.Lock()
	item := s.items[len(s.items) - 1]
	s.items = s.items[:len(s.items) - 1]
	s.lock.Unlock()
	return &item
}

/**
 * 栈的大小
 */
func (s *Stack)Size() int {
	return len(s.items)
}
