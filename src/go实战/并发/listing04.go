package main

import (
	"sync"
	"runtime"
	"fmt"
)

var wg sync.WaitGroup

func main(){
	runtime.GOMAXPROCS(1)
	wg.Add(2)

	go printPrime("a")
	go printPrime("b")

	wg.Wait()
}

func printPrime(prefix string){
	defer wg.Done()

next:
	for outer:=2;outer<5000;outer++{
		for inner := 2;inner<outer;inner++{
			if outer % inner == 0{
				continue next
			}
		}
		fmt.Printf("%s %d",prefix,outer)
	}
}
