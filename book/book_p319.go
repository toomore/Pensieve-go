package main

import "fmt"

type Job struct {
	name string
	desc string
}

func main() {
	joblist := []Job{{"Toomore", "desc"}, {"Toomore2", "desc2"}}
	jobs := make(chan Job)
	done := make(chan bool, len(joblist))
	fmt.Println(jobs, done)

	go func() {
		for _, job := range joblist {
			jobs <- job
		}
		close(jobs)
	}()

	go func() {
		for job := range jobs {
			fmt.Println(job)
			done <- true
		}
	}()

	for i := 0; i < len(joblist); i++ {
		<-done
	}
}
