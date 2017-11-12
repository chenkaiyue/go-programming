package main

import "fmt"

func main(){
	summer := months[6:9]

	for _,s := range summer{
		for _,q := range Q{
			if s == q{
				fmt.Printf("%s",s)
			}
		}
	}
}

//反转slice中的元素
func reverse(s []int){
	for i,j:=0,len(s)-1;i<j;i,j = i+1,j-1{
		s[i],s[j] = s[j],s[i]
	}
}
a := [...]int{0,1,2,3,4,5}
reverse(a[:])
fmt.Println(a)


s := []int{0,1,2,3,4,5}
reverse(s[:2])
reverse(s[2:])
reverse(s)
fmt.Println(s)


func equal(x,y []string) bool{
	if len(x) != len(y){
		return false
	}
	for i := range x{
		if x[i] != y[i]{
			return false
		}
	}
	return true
}


var s []int
s = nil
s = []int(nil)
s = []int{}

make([]int,len,cap)


----------------------append----------------------
func main(){
	var runes []rune
	for _,r := range "hello,world"{
		runes = append(runes,r)
	}
}
func appendInt(x []int,y int) []int{
	var z []int;
	zlen := len(x)+1
	if zlen <= cap(x){
		z = x[:zlen]
	}else{
		zcap := zlen
		if zcap < 2*len(x){
			zcap = 2 * len(x)
		}
		z = make([]int,zlen,zcap)
		copy(z,x)
	}
	z[len(x)] = y
	return z
}

----slice就地修改--------------
func nonempty(strings []string) []string{
	i := 0
	for _,s := range(strings){
		if s != ""{
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
func nonempty2(strings []string) []string{
	out := strings[:0]
	for _,s := range(strings){
		if s != ""{
			out = append(out,s)
		}
	}
	return out
}

-------------------栈的模拟------------------
func main(){
	stack = append(stack,v)
	top := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
}

func remove(slice []int ,i int) []int{
	copy(slice[i:],slice[i+1:])
	return slice[:len(slice)-1]
}

func remove(slice []int, i int) []int{
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}