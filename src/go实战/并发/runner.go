package main

import (
	"os"
	"time"
	"errors"
	"os/signal"
	"github.com/golang/net/html/atom"
)

type Runner struct{
	interrupt chan os.Signal
	complete chan error
	timeout <- chan time.Time
	tasks []func(int)
}

var ErrTimeOut = errors.New("reveived timeout")

var ErrInterrupt = errors.New("reveived interrupt")

func New(d time.Duration) *Runner{
	return &Runner{
		interrupt:make(chan os.Signal,1),
		complete:make(chan error)
		timeout:time.After(d),
	}
}

func (r *Runner) Add(tasks ...func(int)){
	r.tasks = append(r.tasks,tasks...)
}

func (r *Runner) Start() error{
	signal.Notify(r.interrupt,os.Interrupt)

	go func(){
		r.complete<-r.run()
	}()

	select{
	case err:= <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeOut
	}
}

func (r *Runner) run() error{
	for id,task := range(r.tasks){
		if r.gotInterrupt(){
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (r *Runner) gotInterrupt() bool{
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}