package main

import (
	"fmt"
)

var cache [][]int
func find(arr []int, target int) int {
l := len(arr)
if l == 0 {
return -1
}
if l == 1 {
if arr[0] == target {
return target
} else {
return -1
}
}
findSubArr(arr)
// 查找
for i := 0; i < len(cache); i++ {
if len(cache[i]) == 1 {
if target == cache[i][0] {
return target
}
} else {
if target == cache[i][0] {
return cache[i][1]
} else if target == cache[i][1] {
return cache[i][0]
}
}
}
return -1
}

func findSubArr(arr []int) {
if len(arr) <= 2 {
cache = append(cache, arr)
return
}
// 拆分
i := 0
var odd []int
var even []int
for i < len(arr) {
if i % 2 == 0 {
even = append(even, arr[i])
} else {
odd = append(odd, arr[i])
}
i++
}
findSubArr(odd)
findSubArr(even)
}

func main() {
var arr = []int{1,2,3,4,5,6,7,8,9}
res := find(arr, 1)
fmt.Printf("res:%d", res);
}