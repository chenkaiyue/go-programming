package main

import (
	"log"
	"time"
	"github.com/goinaction/code/chapter7/patterns/work"
	"sync"
)

var names=[]string{
	"steve",
	"bob",
	"mary",
}
type namePrinter struct{
	name string
}

func (m *namePrinter) Task(){
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main(){
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(2*len(names))

	for i:=0;i<2;i++{
		for _,name := range names{
			np := namePrinter{
				name:name,
			}
			go func(){
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	p.Shutdown()
}
