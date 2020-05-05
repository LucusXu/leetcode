package alltest

import (
	"testing"
	"fmt"
	"sync"
	"unsafe"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	var count = 1
	var mu sync.Mutex
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(i int) {
			defer func() {
				mu.Unlock()
			}()
			// fmt.Println(i)
			mu.Lock()
			count++
			wg.Done()
		}(i)
	}
	wg.Wait()
	t.Logf("count= %d", count)
}

type Singleton struct {

}

var singleInstance *Singleton
var once sync.Once

func GetSingletonObj() *Singleton{
	once.Do(func() {
		fmt.Println("ccc")
		singleInstance = &Singleton{}
	})
	return singleInstance
}

func TestGetSingle(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0;i < 10;i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("%d\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}