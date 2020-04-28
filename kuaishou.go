package main

import (
	"fmt"
	"sync"
	"time"
	"os"
	"os/signal"
	"syscall"
)

func insertCh(ch chan int, wg *sync.WaitGroup, val int) {
	defer wg.Done()
	ch<-val
}


func out(val int) {
	fmt.Println("%d", val)
}

func insertCh1(ch chan int,  val int) {
	ch<-val
}

func main() {
	ch := make(chan int, 1000)
	// var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		// wg.Add(1)
		go insertCh1(ch, i)
		// go insertCh(ch, &wg, i)
	}
	// wg.Wait()
	t := time.NewTimer(time.Millisecond * 20)
	for {
		select {
		case <-t.C:
			return
		case v := <-ch:
			// fmt.Println("%d", v)
			go out(v)
			break
		default:
			break
		}
	}

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGILL, syscall.SIGINT)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGTERM, syscall.SIGILL, syscall.SIGINT:
				os.Exit(0)
			default:
				break
			}
		}
	}()

}
