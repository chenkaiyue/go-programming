package main

import (
	"sync"
	"runtime"
)

func main(){
	var wg sync.WaitGroup
	//runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GOMAXPROCS(2)
	wg.Add(2)

	go func(){
		defer wg.Done()

		for count:=0;count<50;count++{
			for char:='a';char<'a'+26;char++{
				println("1111111%c",char)
			}
		}
	}()

	go func(){
		defer wg.Done()

		for count:=0;count<50;count++{
			for char:='A';char<'A'+26;char++{
				println("222222%c",char)
			}
		}
	}()

	wg.Wait()

}
