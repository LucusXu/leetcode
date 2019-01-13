package linkList
// 单链表右移k位

import (
)

func Right(head *LinkNode, k int) {
	if head == nil {
		return
	}
	var p = new(LinkNode)
	p = head
	n := 0
	for ; p != nil ; {
		n++
		p = p.next
	}
	// 防止k比n大
	k = n % k
	var fast = new(LinkNode)
	var slow = new(LinkNode)
	i := 0
	// 快指针先走k步
	for ; fast != nil; {
		i++
		fast = fast.next
		if i == k {
			break
		}
	}
	for ; fast.next != nil; {
		fast = fast.next
		slow = slow.next
	}
	fast.next = head
	head = slow.next
	slow.next = nil
}
