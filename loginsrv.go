package main
import (
	"net/http"
	"golang.org/x/net/websocket"
	"chess/fw"
	"chess/login"
	log "github.com/lkj01010/log"
	"chess/cfg"
)



func main() {
	server, err := login.NewServer()
	if err != nil {
		panic("new server failed")
	}
	defer func() {
		server.Close()
	}()

	serve := func(ws *websocket.Conn) {
		if err := server.Serve(fw.NewWsReadWriter(ws)); err != nil {
			log.Error(err.Error())
		}
		log.Infof("new agent comes, agent count=%v", server.AgentCount())
	}

	http.Handle("/", websocket.Handler(serve))
	http.ListenAndServe(":" + cfg.LoginPort, nil)
}