package linkList
// 每隔k个元素翻转单链表

import (
	"fmt"
)

func Rollback(array []int, length, k int) *LinkNode {
	if array == nil || length <= 0 {
		return nil
	}
	i := 0
	head := new(LinkNode)
	head.value = array[i]
	p := head
	fmt.Printf("%v,", head.value)
	for i = 1; i < length; i++ {
		tmp := new(LinkNode)
		tmp.value = array[i]
		p.next = tmp
		p = tmp
		fmt.Printf("%v,", array[i])
	}
	fmt.Printf("\n")

	q := head
	for ;; {
		if q == nil {
			break
		}
		fmt.Printf("%v,", q.value)
		q = q.next
	}
	fmt.Printf("\n")

	head = rollback(head, k)
	q = head
	for ;; {
		if q.next == nil {
			break
		}
		fmt.Printf("%v,", q.value)
		q = q.next
	}
	fmt.Printf("\n")
	return head
}

func rollback(head *LinkNode, k int) *LinkNode {
	if k <= 0 {
		return nil
	}
	prev := new(LinkNode)
	next := new(LinkNode)
	cur := head
	for idx := 0; idx < k; idx++ {
		if cur == nil {
			break
		}
		next = cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	if cur != nil {
		head.next = rollback(next, k)
	}
	return prev
}
