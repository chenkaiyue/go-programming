package main

import (
	"unicode"
	"fmt"
)

func main(){
	slice := make([]int,3,5)

	slice:=[]string{"1","2","3","4","5"}
	slice:=[]int{1,2,3}

	slice := make([]int ,0)

	slice := []int{}


	slice :=[]int {10,20,30,40,50}
	newSlice := slice[1:3]//容量是4，长度为2，容量是剩下的所有的空间
	newSlice = append(newSlice,60) //容量现在是4，长度为3

	newSlice = slice[2:3:4]

	//设置长度和容量一样
	newSlice = slice[2:3:3]
	newSlice = append(newSlice,60)

	//合并切片
	s1 := []int{1,2}
	s2 := []int{3,4}
	fmt.Printf("%v",append(s1,s2...))

	//迭代切片
	for index,value := range slice{
		fmt.Printf("%d:%d",index,value)
	}

	for index := 2;index < len(slice);index++{
		fmt.Printf("%d:%d",index,slice[index])
	}

	//多维切片
	slice := [][]int{{1,2},{10}}
	slice[0] = append(slice[0],20)

	//函数传递切片
	slice := make([]int,1e6)
	slice = foo(slice)
	func foo(slice []int)[]int{

	}
}

