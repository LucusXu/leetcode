package linkList
// 随机从链表中返回一个节点的值
// 如果链表长度未知，则选择用蓄水池算法

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max - min)
	randNum = randNum + min
	fmt.Printf("rand is %v\n", randNum)
	return randNum
}

func Search(head *LinkNode) {
	if head == nil {
		return
	}
	var p = new(LinkNode)
	p = head.next
	cnt := 2
	var res = 1
	for ; p != nil ; {
		res = GenerateRangeNum(1, cnt)
		if res == 1 {
			fmt.Printf("%d\n", p.value)
		}
		cnt++
		p = p.next
	}
	fmt.Printf("%d\n", res)
}
