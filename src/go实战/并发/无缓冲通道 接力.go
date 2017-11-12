package main

import (
	"sync"
	"fmt"
	"time"
)

var wg sync.WaitGroup

func main(){
	baton := make(chan int)

	wg.Add(1)

	go Runner(baton)

	baton <- 1

	wg.Wait()
}

func Runner(baton chan int){
	//defer wg.Done()
	var newrunner int
	runner:= <- baton
	fmt.Printf("%d is running\n",runner)

	if runner!=4{
		newrunner = runner + 1
		fmt.Printf("%d wait\n",newrunner)
		go Runner(baton)
	}

	time.Sleep(100*time.Millisecond)

	if runner == 4{
		fmt.Printf("%d finisher,race over\n",runner)
		wg.Done()
		return
	}

	fmt.Printf("%d to %d\n",runner,newrunner)

	baton<-newrunner
}