package linkList
// 判断单链表是否有环

import (
	"fmt"
)

func HasCircle(head *LinkNode) bool {
	if head == nil || head.next == nil {
		return false
	}

	fast := new(LinkNode)
	slow := new(LinkNode)

	for ; fast.next.next != nil; {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}
	fmt.Println("\n")
	return false
}


// 合并两个有序单链表
func NewMerge(head1, head2 *LinkNode) *LinkNode{
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	head := new(LinkNode)
	// interface类型
	p1 := head1
	p2 := head2
	if (head1.value).(int) < (head2.value).(int) {
		head = head1
		p1 = head1.next
	} else {
		head = head2
		p2 = head2.next
	}
	p := head
	for ;p1 != nil && p2 != nil; {
		if (p1.value).(int) <= (p2.value).(int) {
			p.next = p1
			p = p1
			p1 = p1.next
		} else {
			p.next = p2
			p = p2
			p2 = p2.next
		}
	}
	if p1 != nil {
		p.next = p1
	}
	if p2 != nil {
		p.next = p2
	}
	return head
}

// 删除链表倒数第k个元素
func DelK(head *LinkNode, k int) *LinkNode {
	if head == nil {
		return nil
	}
	p := new(LinkNode)
	len := 0
	for ; p != nil; {
		len++
		p = p.next
	}
	if len < k {
		return head
	}

	if len == k {
		p = head
		// del p
		head = head.next
	} else {
		// 快慢指针
		fast := head
		slow := head
		for ; fast.next != nil; {
			fast = fast.next
			slow = slow.next
		}
		// 删除slow的next
		p = slow.next
		// del p
		slow.next = slow.next.next
	}
	return head
}

// 求链表的中间节点
func MidNode(head *LinkNode) *LinkNode {
	if head == nil {
		return nil
	}
	p := new(LinkNode)
	len := 0
	for ; p != nil; {
		len++
		p = p.next
	}
	// 就一个或者两个节点就是第一个
	if len == 1 || len == 2{
		return head
	}
	// 快慢指针
	fast := new(LinkNode)
	slow := new(LinkNode)
	mid := len >> 1
	for ;mid > 0 && fast.next != nil; {
		fast = fast.next
	}
	for ; fast.next != nil; {
		slow = slow.next
		fast = fast.next
	}
	return slow
}
