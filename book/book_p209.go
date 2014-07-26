package main

import "fmt"
import "math/rand"
import "time"

func main() {
	channels := make([]chan bool, 6)
	for i := range channels {
		channels[i] = make(chan bool)
	}
	go func() {
		rand.Seed(time.Now().UTC().UnixNano())
		channels[rand.Intn(6)] <- true
	}()
	var x int
	select {
	case <-channels[0]:
		x = 0
	case <-channels[1]:
		x = 1
	case <-channels[2]:
		x = 2
	case <-channels[3]:
		x = 3
	case <-channels[4]:
		x = 4
	case <-channels[5]:
		x = 5
	}
	fmt.Println(x)
	//fmt.Println(<-channels[0]) //fail
}
