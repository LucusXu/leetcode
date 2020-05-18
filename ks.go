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


func longestPalindrome(str string) string {
	if len(str) < 1 {
		return ""
	}
	var start, end = 0, 0
	byteStr := []byte(str)
	for i := 0; i < len(str); i++ {
		len1 := expandCenter(byteStr, i, i)
		len2 := expandCenter(byteStr, i, i + 1)
		tmp := max(len1, len2)
		if tmp > end - start {
			start = i - (tmp - 1) / 2
			end = i + tmp / 2
		}
	}
	return string(byteStr[start:end + 1])
}

func expandCenter(str []byte, left, right int) int {
	l := left
	r := right
	for l >= 0 && r < len(str) && str[l] == str[r] {
		l--
		r++
	}
	return r - l - 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	var str = "ababa"
	fmt.Printf("%d\n", longestPalind(str))
	fmt.Printf("%s", longestPalindrome(str))

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