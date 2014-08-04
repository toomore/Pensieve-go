package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
)

func HttpGet(url string, done chan []byte, index int) (data []byte) {
	fmt.Println(index, url, time.Now())
	result, _ := http.Get(url)
	defer result.Body.Close()
	data, _ = ioutil.ReadAll(result.Body)
	done <- data
	fmt.Println(index, url, time.Now())
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

	go func() {
		fmt.Println("123")
		for index, url := range urls {
			result := HttpGet(url, done, index)
			fmt.Println(index, url, len(result))
		}
	}()

	for i := 0; i < len(urls); i++ {
		<-done
	}
}
