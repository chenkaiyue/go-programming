package main

import (
	"fmt"
	"math/rand"
)

func main(){
	//初始化
	dict := make(map[string]int)
	dict := map[string]string{"read":"sdfsa","orange":"12312"}
	dict := map[int][]int{}


	//声明没有实际创建之前不可以赋值
	var colors map[string]string
	colors["red"] = "sdfas" //error

	//判断键是否存在
	value,exists = colors["blue"]
	if exists{
		fmt.Println(value)
	}

	value = colors["blue"]
	if value != ""{
		fmt.Println(value)
	}

	for key,value := range colors{
		fmt.Println(key,value)
	}


	//从map中删除对应的键
	delete(colors,"coral")

	//函数中传递map
	func removeColor(colors map[string]string,key string){
		delete(colors,key)
	}

}