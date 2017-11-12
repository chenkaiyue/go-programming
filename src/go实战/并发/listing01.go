package main

import (
	"runtime"
	"sync"
	"fmt"
)

func main(){
	//分配逻辑处理器给调度器
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("start goroutines")

	go func(){
		defer wg.Done()

		for count:=0;count<3;count++{
			for char:='a';char < 'a'+26;char++{
				fmt.Printf("%c",char)
			}
		}
	}()

	go func(){
		defer wg.Done()

		for count:=0;count<3;count++{
			for char := 'A';char <'A'+26;char++{
				fmt.Printf("%c",char)
			}
		}
	}()

	fmt.Println("wait to finish")

	wg.Wait()

	fmt.Println("\nterminating program")
}