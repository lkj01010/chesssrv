package main
import (
	"net/http"
	"golang.org/x/net/websocket"
	"strconv"
	"chess/fw"
	"fmt"
	"chess/data"
	"chess/login"
)
//func

func serve(ws *websocket.Conn) {
	connCnt ++
	fmt.Printf("agent come, access cnt=%s\n", strconv.Itoa(connCnt))

	agent := fw.NewIpcAgent(&login.LoginServer{}, fw.NewWsReadWriter(ws))
	agent.Serve()
}
var (
	connCnt = 0
	dUser *data.User
)
func onInit() {
	dUser = new(data.User)
	dUser.Init()
}

func onExit() {
	dUser.Exit()
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

	http.Handle("/", websocket.Handler(serve))
	http.ListenAndServe(":8000", nil)
}