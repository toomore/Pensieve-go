package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"sync"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func DoGet(baseURL *url.URL, path string, hostnames []string) {
	//in := bufio.NewReader(os.Stdin)
	//path, _ := in.ReadString('\n')
	//path = strings.Replace(path, "\n", "", -1)

	log.Println("baseURL:", baseURL)
	log.Println("path:", path)
	log.Println("hostnames:", hostnames)

	baseURL.Path = path
	var basehost = baseURL.Host

	var wg sync.WaitGroup
	wg.Add(len(hostnames))

	done := make(chan string, len(hostnames))
	//done := make(chan string, runtime.NumCPU())
	for _, hostname := range hostnames {
		baseURL.Host = hostname + "." + basehost

		go func(URLpath string) {
			defer wg.Done()
			runtime.Gosched()
			var result string
			start := time.Now()
			result = fmt.Sprintf("[%s] [start] %s\n", URLpath, start)
			resp, _ := http.Get(URLpath)
			result += fmt.Sprintf("[%s] [status] %s\n", URLpath, resp.Status)
			result += fmt.Sprintf("[%s] [done] %s\n", URLpath, time.Now().Sub(start))
			done <- result
		}(baseURL.String())
	}

	go func() {
		defer wg.Done()
		for msg := range done {
			fmt.Println(msg)
		}
	}()
	wg.Wait()
	fmt.Println("All Done!")
}

var path = flag.String("path", "/", "Path.")
var baseURLStr = flag.String("base", "http://google.com", "Base url.")
var hosts = flag.String("hosts", "www,www", "Hostname.")

func main() {
	flag.Parse()
	baseURL, err := url.Parse(*baseURLStr)
	if err != nil {
		log.Fatal(err)
	}
	hostnames := strings.Split(*hosts, ",")
	DoGet(baseURL, *path, hostnames)
}
