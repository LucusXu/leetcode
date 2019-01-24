package linkList
// 约瑟夫环问题

import (
	// "fmt"
)

func createCircle(n int) *LinkNode {
	if n <= 0 {
		return nil
	}
	head := new(LinkNode)
	head.value = 1
	p := head
	for i := 2; i <= n; i++ {
		q := new(LinkNode)
		q.value = i
		p.next = q
		p = q
	}
	p.next = head
	return head
}

/**
 * 利用循环链表实现约瑟夫环
 */
func JCircle(n, k int) int {
	if n <= 0 || k <= 0 {
		return 0
	}
	head := createCircle(n)
	for ;; {
		for i := 0; i < k; i++ {
			if i == k - 1 {
				head.next = head.next.next
				head = head.next
				break
			}
			if head.next == head {
				return (head.value).(int)
			}
			head = head.next
		}
	}
	return 0
}
