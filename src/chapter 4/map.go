package main

import (
	"fmt"
	"sort"
	"bufio"
	"os"
	"math"
)

func main(){
	ages := make(map[string]int)
	ages := make(map[string]int){
		"alice":31,
		"chaerlie":34,
	}

	delete(ages["alic"])

	for name,age := range ages{
		fmt.Printf("%s \t %d",name,age)
	}

	var names []string
	for name,age := range ages{
		names = append(names,name)
	}
	sort.Strings(names)
	for _,name := range names{
		fmt.Printf("%s %d",name,ages[name])
	}


	names := make([]string ,0,len(ages))

	var ages map[string]int
	ages == nil


	//键是否存在
	age,ok := ages["bob"]

	if age,ok := ages["bob"];!ok{
		//不存在bob键
	}

}

func equal(x,y map[string]int) bool{
	if len(x) != len(y){
		return false
	}
	for k,xv := range x{
		if yv,ok:=y[k]; !ok || yv != xv{
			return false
		}
	}
	return true
}


func main(){
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		line:=input.Text()
		if !seen[line]{
			seen[line] = true
			fmt.Printf("%s",line)
		}
	}
	if err:=input.Err();err != nil{
		fmt.Fprintf(os.Stderr,"%v",err)
	}
}

var m = make(map[string]int)
func Add(list []string){
	m[k(list)]++
}
func k(list []string) string{
	return fmt.Sprintf("%q",list)
}
func Count(list []string) int{
	return m[k(list)]
}

//--------------------
//graph.go
var graph = make(map[string]map[string]bool)
func addEdge(from,to string){
	edges := graph[from]
	if edges ==nil {
		edges=make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true;
}
func hasEdge(from,to string)bool{
	return graph[from][to]
}