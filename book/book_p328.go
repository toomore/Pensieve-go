package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
)

func HttpGet(url string, done chan []byte, index int) (data []byte) {
	fmt.Println("start", index, url, time.Now())
	result, _ := http.Get(url)
	defer result.Body.Close()
	data, _ = ioutil.ReadAll(result.Body)
	done <- data
	fmt.Println("End", index, url, time.Now())
	return
}

func main() {
	worker := runtime.NumCPU()
	runtime.GOMAXPROCS(worker)
	done := make(chan []byte, worker)
	var urls []string
	urls = append(urls, "http://httpbin.org/get")
	urls = append(urls, "http://google.com/")
	urls = append(urls, "http://www.pinkoi.com")
	urls = append(urls, "http://toomore.net/")

	for index, url := range urls {
		go HttpGet(url, done, index)
	}

	for i := 0; i < len(urls); i++ {
		<-done
	}
	close(done)
}
