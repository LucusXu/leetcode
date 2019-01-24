package linkList

import (
	"fmt"
)

func Lru(array []int, length, n int) *LinkNode {
	if array == nil || length <= 0 || n <= 0 {
		return nil
	}
	head := new(LinkNode)

	record := make(map[int]bool, 0)
	for _, i := range array {
		var tmp = new(LinkNode)
		tmp.value = i
		tmp.next = head
		head = tmp
		p := head

		// 满了
		if len(record) >= n {
			// 命中
			if record[i] == true {
				for ; p.next != nil ; {
					if p.next.value == i {
						p.next = p.next.next
						break
					}
					p = p.next
				}
			} else {
				record[i] = true
				for ; p.next != nil ; {
					if p.next.next == nil {
						record[p.next.value] = false
						p.next = nil
						break
					}
					p = p.next
				}
			}
		} else {
			record[i] = true
		}
	}
	fmt.Printf("%v", record)
	return head
}

func CreateLink(str string) *LinkNode {
	if len(str) <= 0 {
		return nil
	}
	head := new(LinkNode)
	head.value = str[0]
	head.next = nil
	if len(str) > 1 {
		i := 1
		p := head
		for ; i < len(str); i++ {
			tmp := new(LinkNode)
			tmp.value = str[i]
			p.next = tmp
			p = p.next
		}
		p.next = nil
	}

	p := head
	for ;p != nil; {
		fmt.Println("%s", p.value)
		p = p.next
	}
	fmt.Println("\n")
	return head
}

func Check(head *LinkNode) bool {
	if head == nil {
		return true
	}
	p := head
	// 链表长度
	len := 0

	for ; p != nil; {
		p = p.next
		len++
	}
	fmt.Println("%d", len)
	// 找到中间节点
	m := len >> 1
	mid := head
	for ; m > 0 && mid != nil; {
		mid = mid.next
		m--
	}

	// 翻转右边链表
	prev := mid
	cur := mid.next
	for ; cur.next != nil; {
		next := cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	// cur是另一端的head
	// 判断
	p = head
	q := cur
	var checkRet bool = true
	for ; p != q; {
		if p.value != q.value {
			checkRet = false
			break
		}
		p = p.next
		q = q.next
	}
	// 恢复后半段链表
	prev = nil
	for ; cur != mid; {
		next := cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	return checkRet
}
