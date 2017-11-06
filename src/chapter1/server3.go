package main

import (
	"net/http"
	"fmt"
	"log"
)

func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"%s %s %s\n",r.Method,r.URL,r.Proto)
	for k,v := range r.Header{
		fmt.Printf(w,"Header[%q] = %q\n",k,v)
	}
	fmt.Printf(w,"host=%q\n",r.Host)
	fmt.Printf(w,"remoteaddr=%q\n",r.RemoteAddr)
	if err := r.ParseForm(); err != nil{
		log.Print(err)
	}
	for k,v := range r.Form{
		fmt.Printf(w,"form[%q] = %q \n",k,v)
	}
}
