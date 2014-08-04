package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
)

func HttpGet(url string, done chan []byte, index int) (data []byte) {
	result, _ := http.Get(url)
	defer result.Body.Close()
	data, _ = ioutil.ReadAll(result.Body)
	fmt.Println(index, url)
	done <- data
	return
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	done := make(chan []byte, 5)
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
		//fmt.Println(<-done)
	}
	close(done)
}
