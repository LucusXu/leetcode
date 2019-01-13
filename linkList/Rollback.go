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

// 翻转链表
func Reverse(head, left *LinkNode) *LinkNode {
	if head == nil || left == nil {
		return nil
	}
	// 只有一个节点了
	if left.next.next == nil {
		return head
	}
	cur := left
	prev := left
	p := left.next
	for ; p.next != nil; {
		prev = p
		p = p.next
	}
	prev.next = nil
	p.next = cur.next
	cur.next = p
	// 递归
	head = Reverse(head, p)
	return head
}

// 在单链表上对折
func Fold(head *LinkNode) *LinkNode {
	if head == nil {
		return nil
	}
	// 1.先找到中点, 快慢指针
	mid := Mid(head)
	// 2.翻转右边链表
	head = Reverse(head, mid)
	// 3.归并左右半个链表
	head = Merge(head, mid)
	return head
}

// 归并链表左右各半
func Merge(head, mid *LinkNode) *LinkNode {
	if head == nil || mid == nil {
		return nil
	}

	p := head
	q := mid.next

	for ; q.next != nil; {
		mid.next = q.next
		q.next = p.next
		p.next = q
		p = p.next.next
		q = mid.next
	}
	return head
}

// 在单链表上对折
func Mid(head *LinkNode) *LinkNode {
	if head == nil {
		return nil
	}
	p := head
	var len = 0;
	for ; p.next != nil; {
		len++
		p = p.next
	}
	var min = len / 2
	// 快慢指针
	fast := head
	var i = 1
	for ;; {
		if i >= min {
			break
		}
		i++
		fast = fast.next
	}
	return fast
}
