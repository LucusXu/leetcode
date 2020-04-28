package main

import (
	"fmt"
)

func main() {
	str := "你好"
	fmt.Println("%d", len(str), len([]byte(str)), len([]rune(str)))
}
