package main

import "fmt"

type notifier interface{
	notify()
}


type user struct{
	name string
	email string
}

type admin struct{
	user
	level string
}

func (u *user) notify(){
	fmt.Printf("user email %s:%s",u.name,u.email)
}

func (a *admin) notify(){
	fmt.Printf("admin email %s: %s",u.name,u.email)
}

func main() {
	ad := admin{
		user: user{
			name:  "john",
			email: "@.com",
		},
		level: "super",
	}
	sendNotification(&ad)
	ad.user.notify()
	ad.notify()
}

func sendNotification(n notifier){
	n.notify()
}