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

// 排序单链表（奇数位升序，偶数位降序）
func Sort(head *LinkNode) *LinkNode {
	if head == nil {
		return nil
	}
	// 1.先拆分链表，奇数一个，偶数一个
	head1, head2 := Separate(head)
	// 2.翻转降序链表
	head2 = FullReverse(head2)
	// 3.归并左右半个链表
	head = MergeSort(head1, head2)
	return head
}

func MergeSort(head1, head2 *LinkNode) *LinkNode {
	if head1 == nil || head2 == nil {
		return nil
	}
	p := head1
	q := head2
	head := new(LinkNode)
	if p.value <= q.value {
		head = head1
		p = head1.next
	} else {
		head = head2
		q = head2.next
	}
	x := head
	for ; p != nil && q != nil; {
		if p.value <= q.value {
			x.next = p
			x = p
			p = p.next
		} else {
			x.next = q
			x = q
			q = q.next
		}
	}
	if p != nil {
		x.next = p
	}
	if q != nil {
		x.next = q
	}
	return head
}

func Separate(head *LinkNode) (*LinkNode, *LinkNode) {
	if head == nil {
		return nil, nil
	}
	var i int = 1
	head1 := head
	head2 := head
	p := head
	p1 := head1
	p2 := head2

	for ;p != nil; {
		// 奇数
		if i % 2 == 1 {
			p1.next = p
		} else {
			p2.next = p
		}
		i++
		p = p.next
	}
	return head1, head2
}

// 翻转链表
func FullReverse(head *LinkNode) *LinkNode {
	if head == nil {
		return nil
	}
	// 只有一个节点了
	if head.next == nil {
		return head
	}

	prev := new(LinkNode)
	cur := head
	p := head.next
	for ; p.next != nil; {
		cur.next = prev
		prev = cur
		cur = p
	}
	return cur
}




