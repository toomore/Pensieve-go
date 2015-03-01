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

func DoGet(baseURL *url.URL, path string, hostnames []string) {

	log.Println("baseURL:", baseURL)
	log.Println("path:", path)
	log.Println("hostnames:", hostnames)
	log.Println("CPU Nums:", *nCPU)

	baseURL.Path = path

	var wg sync.WaitGroup
	wg.Add(len(hostnames) * 2)

	first := time.Now()
	done := make(chan []byte, len(hostnames))
	for _, hostname := range hostnames {
		var dobaseURL = *baseURL
		dobaseURL.Host = hostname + "." + dobaseURL.Host

		go func(URLpath string) {
			defer wg.Done()
			runtime.Gosched()
			var result []byte
			start := time.Now()
			result = append(result, fmt.Sprintf("[%s] [start] %s\n", URLpath, start)...)
			resp, _ := http.Get(URLpath)
			result = append(result, fmt.Sprintf("[%s] [status] %s\n", URLpath, resp.Status)...)

			var doneNow = time.Now()
			var doneTime = doneNow.Sub(start)
			var doneFromInit = doneNow.Sub(first)

			result = append(result, fmt.Sprintf("[%s] [done] %s %s (%s)\n", URLpath, doneTime, doneFromInit, doneFromInit-doneTime)...)

			done <- result
		}(dobaseURL.String())
	}

	go func() {
		for msg := range done {
			fmt.Printf("%s", msg)
			wg.Done()
		}
	}()
	wg.Wait()
	fmt.Println("All Done!")
}

var path = flag.String("path", "/", "Path.")
var baseURLStr = flag.String("base", "http://google.com", "Base url.")
var hosts = flag.String("hosts", "docs,www", "Hostname.")
var nCPU = flag.Int("cpus", runtime.NumCPU(), "NumCPU.")

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(*nCPU)
	baseURL, err := url.Parse(*baseURLStr)
	if err != nil {
		log.Fatal(err)
	}
	hostnames := strings.Split(*hosts, ",")
	DoGet(baseURL, *path, hostnames)
}
