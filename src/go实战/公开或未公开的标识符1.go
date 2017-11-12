package main

import (
	"fmt"
	"go实战/counters"
)

func main(){
	counter := counters.New(10)
	fmt.Printf("counter:%d",counter)
}
