package main

import (
	"fmt"
	"os"
	"log"
)

func newInt() *int{
	return new(int)
}

func newInt() *int{
	var dummy int
	return &dummy
}

p := new(int)
fmt.Println(*p)

func gcd(x,y int) int {
	for y != 0{
		x,y=y,x%y
	}
	return x
}

func fib(n int) int{
	x,y := 0,1
	for i := 0;i<n;i++{
		x,y = y,x+y
	}
	return x
}


隐式赋值
medals := []string{"gold","silver","bronze"}

type Celsius float64


func main(){
	x:="hello"
	for _,x := range x{
		x := x+'A'-'a'
		fmt.Printf("%c",x)
	}

	for i:=0;i<len(x);i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x)
		}
	}
	if f,err := os.Open(fname);err!=nil{
		return err
	}else{
		f.Stat()
		f.Close()
	}
}

var cwd string
func init(){
	var err error
	cwd,err = os.Getwd()
	if err != nil{
		log.Fatalf("%v",err)
	}
}