package main
import (
	"net/http"
	"golang.org/x/net/websocket"
	"chess/fw"
	"chess/login"
	log "github.com/lkj01010/log"
)

func serve(ws *websocket.Conn) {
	if err := server.Serve(fw.NewWsReadWriter(ws)); err != nil {
		log.Error(err.Error())
	}
	log.Infof("new agent comes, agent count=%v", len(server.AgentCount()))
}

var server *login.Server

func main() {
	defer func() {
		server.Close()
	}()

	//	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	//		s := http.Server{Handler: websocket.Handler(wsHandler)}
	//		s.ServeHTTP(w, req)
	//	}
	//
	//	err := http.ListenAndServe(":"+strconv.Itoa(Config.Port), nil)
	//	if err != nil {
	//		panic("Error: " + err.Error())
	//	}
	server = login.NewServer()

	http.Handle("/", websocket.Handler(serve))
	http.ListenAndServe(":8000", nil)
}