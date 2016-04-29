package main

import "log"

func main() {
	//var m = make(map[int]int)
	var m = make(map[int64]int)
	for i := 0; i < 5000000; i++ {
		//var stats runtime.MemStats
		//runtime.GC()
		//runtime.ReadMemStats(&stats)
		log.Println(stats.Alloc)
		m[int64(i)] = 1
	}
}
