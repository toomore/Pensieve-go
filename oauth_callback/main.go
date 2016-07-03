package main

import (
	"io"
	"log"
	"net/http"
)

func callback(w http.ResponseWriter, r *http.Request) {
	log.Printf("%+v", r.URL.Query())
	io.WriteString(w, r.URL.Query().Encode())
}

func main() {
	http.HandleFunc("/callback", callback)
	log.Println("Starting ...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8888", nil))
}
