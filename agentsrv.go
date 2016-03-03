package main
import (
	"net/http"
	"golang.org/x/net/websocket"
	"chess/fw"
	"chess/agent"
	log "github.com/lkj01010/log"
	"chess/cfg"
)

func main() {
	server, err := agent.NewServer()
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
		log.Infof("agent leaves, agent count=%v", server.AgentCount())
	}

	http.Handle("/", websocket.Handler(serve))
	log.Debug("agent server start on:", cfg.AgentPort)
	http.ListenAndServe(":" + cfg.AgentPort, nil)
}