package main

import (
	"io"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

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
	for {
		if n, err := ws.Read(msg); err == nil {
			log.Println(ws.Request().Header["Sec-Websocket-Key"], msg[:n])
			ws.Write(msg[:n])
		} else {
			log.Println(err)
			break
		}
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
