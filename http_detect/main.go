package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

func DoGet(baseURL *url.URL, path string, hostnames []string) {
	llog := log.New(os.Stdout, " - ", 0)

	llog.Println("Date:", time.Now())
	llog.Println("baseURL:", baseURL)
	llog.Println("path:", path)
	llog.Println("hostnames:", hostnames)
	llog.Println("CPU Nums:", *nCPU)
	llog.Printf("NonStop: %d ms \n", *nonstop)

	baseURL.Path = path

	var wg sync.WaitGroup
	wg.Add(len(hostnames))

	first := time.Now()
	done := make(chan []byte, *nCPU)
	defer close(done)
	for _, hostname := range hostnames {
		var result []byte
		var dobaseURL = *baseURL
		dobaseURL.Host = hostname + "." + dobaseURL.Host

		go func(URLpath string) {
			runtime.Gosched()
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
		runtime.Gosched()
		for msg := range done {
			fmt.Printf("%s", msg)
			wg.Done()
		}
	}()
	wg.Wait()
	fmt.Println("All Done!")
}

var (
	path       = flag.String("path", "/", "Path.")
	baseURLStr = flag.String("base", "http://google.com", "Base url.")
	hosts      = flag.String("hosts", "docs,www", "Hostname.")
	nCPU       = flag.Int("cpus", runtime.NumCPU(), "NumCPU.")
	nonstop    = flag.Int64("nonstop", 0, "Non-stop in ms.")
)

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(*nCPU)
	var (
		baseURL *url.URL
		err     error
	)

	if baseURL, err = url.Parse(*baseURLStr); err != nil {
		log.Fatal(err)
	}

	hostnames := strings.Split(*hosts, ",")
	for {
		DoGet(baseURL, *path, hostnames)
		if *nonstop > 0 {
			time.Sleep(time.Duration(*nonstop) * time.Millisecond)
		} else {
			os.Exit(0)
		}
	}
}
