package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

// Jobs for run
type Jobs interface {
	Run()
}

func run(jobs Jobs, sleepTime time.Duration, ch chan interface{}) {
	go func() {
		for {
			fmt.Println(time.Now())
			jobs.Run()
			time.Sleep(sleepTime)
		}
	}()
}

// AddOne struct
type AddOne struct {
	Name string
}

// Run func
func (a *AddOne) Run() {
	fmt.Println("1+1", a.Name)
}

func main() {
	debug.SetGCPercent(-1)
	endChan := make(chan interface{})
	works := &AddOne{Name: "Add One"}
	run(works, time.Second, endChan)
	<-endChan
}
