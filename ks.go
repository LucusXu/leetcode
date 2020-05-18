package main

import (
	"fmt"
)

func longestPalind(str string) int {
	var cache = make(map[byte]int, 0)
	bstr := []byte(str)
	for _, v := range bstr {
		if _, ok := cache[v]; !ok {
			cache[v] = 1
		} else {
			cache[v]++
		}
	}
	var res = 0
	for _, value := range cache {
		res = res + value / 2 * 2
		if res % 2 == 0 && value % 2 == 1 {
			res++
		}
	}
	return res
}


func main() {
	var str = "abccccdd"
	fmt.Printf("%d\n", longestPalind(str))


	var ch = make(chan int, 1000)
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println("i", <-ch)
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			ch<-i
		}
	}()
}