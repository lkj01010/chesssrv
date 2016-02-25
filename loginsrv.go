package main
import (
	"net/http"
	"golang.org/x/net/websocket"
	"strconv"
	"chess/fw"
	"fmt"
	"chess/dao"
	"chess/login"
)
//func

func serve(ws *websocket.Conn) {
	connCnt ++
	fmt.Printf("agent come, access cnt=%s\n", strconv.Itoa(connCnt))

	agent := fw.NewAgent(&login.handler{}, fw.NewWsReadWriter(ws))
	agent.Serve()
}
var (
	connCnt = 0
)
func onInit() {
}

func onExit() {
	dao.Exit()
}

func main() {
	onInit()
	defer func() {
		onExit()
	}()
	//	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	//		s := http.Server{Handler: websocket.Handler(wsHandler)}
	//		s.ServeHTTP(w, req)
	//	}r
	//
	//	err := http.ListenAndServe(":"+strconv.Itoa(Config.Port), nil)
	//	if err != nil {
	//		panic("Error: " + err.Error())
	//	}
	server := login.NewServer()

	http.Handle("/", websocket.Handler(serve))
	http.ListenAndServe(":8000", nil)
}