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
		var (
			dobaseURL = *baseURL
			result    []byte
			status    string
		)
		dobaseURL.Host = hostname + "." + dobaseURL.Host

		go func(URLpath string) {
			runtime.Gosched()
			start := time.Now()
			result = append(result, fmt.Sprintf("[%s] [start] %s\n", URLpath, start)...)

			if resp, err := http.Get(URLpath); err == nil {
				status = resp.Status
			} else {
				status = fmt.Sprintf("[Error] %s", err.Error())
			}

			result = append(result, fmt.Sprintf("[%s] [status] %s\n", URLpath, status)...)

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
	log.Println("All Done!")
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
		baseURL   *url.URL
		err       error
		hostnames []string
		sleeptime time.Duration
	)

	if baseURL, err = url.Parse(*baseURLStr); err != nil {
		log.Fatal(err)
	}

	hostnames = strings.Split(*hosts, ",")
	sleeptime = time.Duration(*nonstop) * time.Millisecond
	if sleeptime > 0 {
		for {
			DoGet(baseURL, *path, hostnames)
			time.Sleep(sleeptime)
		}
	} else {
		DoGet(baseURL, *path, hostnames)
		os.Exit(0)
	}
}
