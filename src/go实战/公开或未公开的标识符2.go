package main

import (
	"go实战/entities"
	"fmt"
)

func main(){
	a := entities.Admin{
		Rights:10,
	}

	a.Email="@.com"
	a.Name ="bill"
	fmt.Printf("%v",a)
}

