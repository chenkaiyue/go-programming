package main

import (
	"sync"
	"runtime"
	"fmt"
)

var(
	counter int
	wg sync.WaitGroup
	mutex sync.Mutex
)

func main(){
	wg.Add(2)

	go inCounter(1)
	go inCounter(2)

	wg.Wait()
	fmt.Printf("%d",counter)
}

func inCounter(id int){
	defer wg.Done()

	for count:=0;count<2;count++{
		mutex.Lock()
		{
			value := counter

			//当前goroutien从线程退出，并放回到队列
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()
	}
}