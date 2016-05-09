package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

var wsPools = make(map[string]*websocket.Conn)

// Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
	//ws.Write([]byte("Toomore"))
	log.Println("1.", ws.Request().Header)
	var msg = make([]byte, 512)
	if n, err := ws.Read(msg); err != nil {
		log.Println(err)
	} else {
		log.Println("2.", ws.RemoteAddr(), n)
	}
	ws.Write([]byte("Done!!"))
	io.Copy(ws, ws)
	log.Println("3. byby")
}

func nonStop(ws *websocket.Conn) {
	var msg = make([]byte, 512)
	log.Println("In: ", ws.Request().Header["Sec-Websocket-Key"])
	wsNo := ws.Request().Header["Sec-Websocket-Key"][0]
	wsPools[wsNo] = ws
	for {
		if n, err := ws.Read(msg); err == nil {
			log.Println(wsNo, msg[:n])
			for _, vws := range wsPools {
				go func(vws *websocket.Conn) {
					vws.Write(msg[:n])
				}(vws)
			}
		} else {
			log.Printf("%s, %+v", err, wsNo)
			ws.Write([]byte("byby"))
			delete(wsPools, wsNo)
			break
		}
	}
	for i := range wsPools {
		fmt.Println("Rest: ", i)
	}
}

// This example demonstrates a trivial echo server.
func main() {
	http.Handle("/echo", websocket.Handler(EchoServer))
	http.Handle("/nonstop", websocket.Handler(nonStop))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
