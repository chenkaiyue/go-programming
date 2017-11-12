package main

import (
	"encoding/xml"
	"image"
)


func main(){
	type Employee struct{
		ID		int
		Name	string
		Position	string
	}
	var dilbert Employee

	position = &dilbert.Position
	*positon = *position + "123123"


	var employee *Employee = &dilbert
	employee.Position += "sdfs"


	/*
	二叉排序
	*/
}

type tree struct{
	value int
	left,right *tree
}

func Sort(values []int){
	var root *tree
	for _,v := range values{
		root = add(root,v)
	}
	appendValues(values[:0],root)
}

func appendValues(values []int,t *tree) []int{
	if t != nil{
		values = appendValues(values,t.left)
		values = append(values,t.value)
		values = appendValues(values,t.right)
	}
	return values
}

func add(t *tree,value int)*tree{
	if t == nil{
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value{
		t.left = add(t.left,value)
	}else{
		t.right = add(t.right,value)
	}
	return t
}





/*
结构体字面量
*/
func main(){
	type Point struct{
		x,y int
	}
	p := Point{1,2}
	q := Point{x:1,y:2}
	pp := &Point{1,2}


	type address struct{
		hostname string
		port int
	}
	hits := make(map[address]int)
	hits[address{"123",22}]++

}

func Scale(p Point,factor int) Point{
	return Point{p.x*factor,p.y*factor}
}



/*
结构体嵌套
*/


func main(){
	type Circle struct{
		Point
		Radius int
	}
	type wheel struct{
		Circle
		Spokes int
	}

	var w = wheel{Circle{Point{1,2},10},20}

}



/*
json
*/

func main(){
	type Movie struct{
		Title string
		yea
	}
}