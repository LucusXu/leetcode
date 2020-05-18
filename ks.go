package main

import (
	"fmt"
)



func main() {
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