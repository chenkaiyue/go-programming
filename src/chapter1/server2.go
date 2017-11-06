package main

import (
	"sync"
	"net/http"
	"log"
	"fmt"
)

var mu sync.Mutex
var count int

func main(){
	http.HandleFunc("/",handler)
	http.HandleFunc("/count",counter)
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

func handler(w http.ResponseWriter,r *http.Request){
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Printf(w,"url.path=%q\n",r.URL.Path)
}

func counter(w http.ResponseWriter,r *http.Request){
	mu.Lock()
	fmt.Printf(w,"count %d\n",count)
	mu.Unlock()
}