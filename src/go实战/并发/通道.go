package main

import (
	"sync"
	"math/rand"
	"time"
	"fmt"
)

var wg sync.WaitGroup

func init(){
	rand.Seed(time.Now().UnixNano())
}

func main(){
	court := make(chan int)

	wg.Add(2)

	go player("n",court)
	go player("d",court)

	court <- 1

	wg.Wait()
}

func player(name string,court chan int){
	defer wg.Done()

	for {
		ball,ok := <-court
		if !ok{
			fmt.Printf("%s won\n",name)
			//wg.Done()
			return
		}

		n := rand.Intn(100)
		if n % 13 == 0{
			fmt.Printf("%s missed\n",name)
			close(court)
			//wg.Done()
			return
		}

		fmt.Printf("%s hit %d\n",name,ball)
		ball++

		court <- ball
	}
}
