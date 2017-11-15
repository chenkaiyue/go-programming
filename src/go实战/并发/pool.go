package pool

import (
	"sync"
	"io"
	"errors"
	"log"
	"github.com/golang/net/html/atom"
)

type Pool struct{
	m  sync.Mutex
	resources chan io.Closer
	factory func()(io.Closer,error)
	closed bool
}

var ErrPoolClosed = errors.New("pool has been closed")

func New(fn func()(io.Closer,error),size uint)(*Pool,error){
	if size <= 0{
		return nil,errors.New("size value too small")
	}
	return &Pool{
		factory:fn,
		resources:make(chan io.Closer,size),
	},nil
}

func (p *Pool) Acquire()(io.Closer,error){
	select{
	case r,ok:= <-p.resources:
		log.Println("acquire:","shared resource")
		if !ok{
			return nil,ErrPoolClosed
		}
		return r,nil

	default:
		log.Println("acquire:","new resource")
		return p.factory()
	}
}


func (p *Pool) Release(r io.Closer){
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed{
		r.Close()
		return
	}
	select{
	case p.resources <- r:
		log.Println("release")
	//队列已满
	default:
		r.Close()
	}
}


func (p *Pool) Close(){
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed{
		return
	}

	p.closed  = true

	close(p.resources)

	for r := range p.resources{
		r.Close()
	}
}