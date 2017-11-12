package main
func main(){
	var array [5]int
	array := [5]int{1,2,3,4,5}

	array:=[...]int{1,2,3}

	array :=[5]int{1:10,2:20}

	var array [5]*int
	array:=[5]*int{0:new(int),1:new(int)}
	*array[0] = 10
	*array[1] = 20


	var array1 [3]*string
	array2 := [3]*string{new(string),new(string),new(string)}
	*array2[0] = "red"
	*array2[1] = "blue"
	*array2[2] = "green"

	array1 = array2



	//二维数组
	var array [4][2]int
	array := [4][2]int{{10,11},{20,21},{20,21},{20,21}}
	array:= [4][2]int{1:{1,2},3:{3,4}}
	array:= [4][2]int{1:{0:1},3:{1:2}}

	var array [1e6]int
	foo(&array)
}

func foo(array *[1e6]int){

}

