package main

import (
	"sync"
	"time"
	"sync/atomic"
	"fmt"
)

var(
	shutdown int64
	wg sync.WaitGroup
)

func main(){
	wg.Add(2)

	go dowork("a")
	go dowork("b")

	time.Sleep(1*time.Second)

	atomic.StoreInt64(&shutdown,1)

	wg.Wait()
}

func dowork(name string){
	defer wg.Done()

	for{
		fmt.Printf("%s work\n",name)
		time.Sleep(250*time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1{
			fmt.Printf("%s down\n",name)
			break
		}
	}
}