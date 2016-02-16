package main

import (
	"net/http"
	"golang.org/x/net/websocket"
	"strconv"
	"chess/fw"
	"fmt"
)

func serve(ws *websocket.Conn) {
	connCnt ++
	fmt.Printf("agent come, access cnt=%s\n", strconv.Itoa(connCnt))

	rw := &fw.WsReadWriter{ws: ws}
	var msg string
	rw.Read(&msg)
}
var (
	connCnt = 0
)
func main() {
	//	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	//		s := http.Server{Handler: websocket.Handler(wsHandler)}
	//		s.ServeHTTP(w, req)
	//	}r
	//
	//	err := http.ListenAndServe(":"+strconv.Itoa(Config.Port), nil)
	//	if err != nil {
	//		panic("Error: " + err.Error())
	//	}

	http.Handle("/", websocket.Handler(serve))
	http.ListenAndServe(":8000", nil)
}