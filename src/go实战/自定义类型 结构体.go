package main

import "fmt"

func main(){
	type user struct{
		name string
		email string
		ext int
		privileged bool
	}
	var bill user
	lisa := user{
		name: "lisa",
		email: ".com"
	}
	lisa := user{"lisa","email.com"}

	type admin struct{
		person user
		level string
	}

	fred := struct{
		person: user{
		name: "lisa",
		email: "email.com",
	}
		level:"super"
	}


	type duration int64
	var d duration
	d = duration(int64(10))

}

//方法
type user struct{
	name string
	email string
}

func (u user) notify(){
	fmt.Println(u.name,u.email)
}

func (u *user) changeEmail(email string){
	u.email = email
}

func main(){
	bill := user{"bill","oldemail"}
	bill.notify()

	lisa := &user{"lisa","loldmail"}
	lisa.notify()

	bill.changeEmail("newmail")
	lisa.changeEmail("lnewmail")
}