package main

import "fmt"

func main(){
	var a [3]int
	for i,v := range a{
		fmt.Printf("%d %d",i,v)
	}
	var q [3]int = [3]int{1,2,3}
	q := [...]int{1,2,3}

	type Cuerrency int
	const(
		USD Cuerrency = iota
		RMB
	)
	symbol := [...]string{USD:"$",RMB:"￥"}
	fmt.Println(RMB,symbol[RMB])

	r:= [...]int{99:1}
}


func zero(ptr *[32]byte){
	for i := range ptr{
		ptr[i]=0
	}
}

func zero1(ptr *[32]byte){
	*ptr = [32]byte{}
}