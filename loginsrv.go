package main
import (
	"net/http"
	"golang.org/x/net/websocket"
	"chess/fw"
	"chess/login"
	log "github.com/lkj01010/log"
)



func main() {
	server := login.NewServer()
	defer func() {
		server.Close()
	}()

	serve := func(ws *websocket.Conn) {
		if err := server.Serve(fw.NewWsReadWriter(ws)); err != nil {
			log.Error(err.Error())
		}
		log.Infof("new agent comes, agent count=%v", len(server.AgentCount()))
	}

	http.Handle("/", websocket.Handler(serve))
	http.ListenAndServe(":8000", nil)
}