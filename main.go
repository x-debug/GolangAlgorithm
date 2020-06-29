package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(20)
	var rwm sync.RWMutex
	data := 0
	for i := 0; i < 10; i++ {
		go func(t int) {
			rwm.RLock()
			defer rwm.RUnlock()
			fmt.Printf("读数据:%v %d\r\n", data, t)
			wg.Done()
			time.Sleep(time.Second * 1)
		}(i)

		go func(t int) {
			rwm.Lock()
			defer rwm.Unlock()
			data += t
			fmt.Printf("写数据:%v %d\r\n", data, t)
			wg.Done()
			time.Sleep(time.Second * 5)
		}(i)
	}

	wg.Wait()
}
