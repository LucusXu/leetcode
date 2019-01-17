package sort
/**
 * 用两个栈实现队列
 */

import (
	// "fmt"
)

type Squeue struct {
	s1 Stack
	s2 Stack
}

/**
 * 入队
 */
func (q *Squeue)Push(v interface{}) {
	q.s1.Push(v)
}

/**
 * 出队
 */
func (q *Squeue)Pop() (v interface{}) {
	size := q.s2.Size()
	if size > 0 {
		v = q.s2.Pop()
		return v
	} else {
		size1 := q.s1.Size()
		if size1 > 0 {
			for i := 0; i < size1;i++ {
				v = q.s1.Pop()
				q.s2.Push(v)
			}
			v = q.s2.Pop()
			return v
		} else {
			return nil
		}
	}
}

// 两个队列实现栈
type QStack struct {
	q1 Queue
	q2 Queue
}

/**
 * 入栈
 */
func (s *QStack)SPush(v interface{}) {
	s.q1.PushBack(v)
}

/**
 * 出栈
 */
func (s *QStack)SPop() (v interface{}) {
	size := s.q1.Len()
	if size == 0 {
		return nil
	}

	for i := 0; i < size - 1; i++ {
		item := s.q1.Pop()
		s.q2.PushBack(item)
	}
	v = s.q1.Pop()
	size2 := s.q2.Len()
	for i := 0; i < size2; i++ {
		item := s.q2.Pop()
		s.q1.PushBack(item)
	}
	return v
}