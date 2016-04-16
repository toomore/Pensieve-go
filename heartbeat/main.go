package main

import (
	"fmt"
	"time"
)

// Jobs for run
type Jobs interface {
	Run()
}

func run(jobs Jobs) {
	for {
		fmt.Println(time.Now())
		jobs.Run()
		time.Sleep(time.Second)
	}
}

type AddOne struct {
	Name string
}

func (a *AddOne) Run() {
	fmt.Println("1+1", a.Name)
}

func main() {
	works := &AddOne{Name: "Add One"}
	run(works)
}
