package linkList
// LRU最久未使用算法

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

