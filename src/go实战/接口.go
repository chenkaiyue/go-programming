package main

import "fmt"

type duration int
func (d *duration) pretty() string{
	return fmt.Sprintf("hello %d",*d)
}


type notifier interface{
	notify()
}

type user1 struct{
	name string
	email string
}

func (u *user1) notify(){
	fmt.Printf("sending %s to %s",u.name,u.email)
}

func main(){
	u := user1{"bill","@com"}
	sendNotification(u)
}

func sendNotification(n notifier){
	n.notify()
}

func main(){
	//a := duration(10)
	//fmt.Println((&a).pretty())
	//fmt.Println(a.pretty())
	addr := &(duration(10))
	fmt.Println(addr.pretty())

}
