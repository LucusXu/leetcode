package sort

import (
	"sync"
	// 队列用链表实现
	"container/list"
)

type Element interface {}

type Queue struct {
	l *list.List
	// 锁
	lock sync.RWMutex
}

/**
 * 创建队列
 */
func New() * Queue {
	return &Queue{l : list.New()}
}

/**
 * 入队列
 */
func (q *Queue)PushBack(v interface{}) {
	if v == nil {
		return
	}
	q.lock.Lock()
	defer q.lock.Unlock()
	q.l.PushBack(v)
}

/**
 * 出队列
 */
func (q *Queue)Pop() (e interface{}) {
	if q.l.Len() == 0 {
		return
	}
	q.lock.Lock()
	defer q.lock.Unlock()
	e = q.l.Front().Value
	q.l.Remove(q.l.Front())
	return e
}

/**
 * 队列长度
 */
func (q *Queue)Len() int {
	len := q.l.Len()
	return len
}
