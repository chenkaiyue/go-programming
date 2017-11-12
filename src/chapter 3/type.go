package main

import (
	"go/ast"
	"fmt"
	"strconv"
	"time"
	"unicode"
)

func main{
	var apples int32 =1
	var oranges int 16= 2
	var compote = int(apples) + int(oranges)
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o \n",o)

	var x complex128 = complex(1,2)
	fmt.Println(real(x))
	fmt.Println(imag(x))
	x := 1+2i

	s:="hello"
	fmt.Println(len(s))
	fmt.Println(s[0])

	for i,r := range  "hello,世界"{
		fmt.Printf("%d %q %d",i,r,r)
	}

	n:=0
	for range s{
		n++
	}


	s := "abc"
	b := []byte(s)
	s2 := string(b)


	x := 123
	fmt.Println(strconv.Itoa(x))
	fmt.Sprintf("%d",x)
	fmt.Sprintf("x=%b",x) //"x=1111011"
	fmt.Println(strconv.FormatInt(int64(x),2))
	x,err := strconv.Atoi("123")
	x,err := strconv.ParseInt("123",10,64)


	const noDelay time.Duration = 0

	const (
		a=1
		b
		c=2
		d
	)
	fmt.Println(a,b,c,d)



	type weekday int
	const(
		Sunday weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	type Flags uint
	const(
		FlagUp Flags = 1 << iota
		FlagBroadcast
		...
	)
}
func HasPrefix(s,prefix string) bool{
	return len(s) > len(prefix) && s[:len(prefix)] == prefix
}

func btoi(b bool) int{
	if b{
		return 1
	}
	else {
		return 0
	}
}



