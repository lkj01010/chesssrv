package main

import (
	"chess/cfg"
	"github.com/lkj01010/log"
	"golang.org/x/net/websocket"
	"net/http"
	"chess/fw"
	"chess/game"
)

//func main() {
//	//register rpc
//	m := game.NewModel()
//	rpc.Register(m.Server)
//	defer func() {
//		m.Exit()
//	}()
//
//	//network
//	serverAddr := cfg.GameAddr()
//	l, e := net.Listen("tcp", serverAddr)
//	if e != nil {
//		log.Fatalf("net.Listen tcp : %v", e)
//	}
//	log.Info("dao RPC server listening on ", serverAddr)
//	rpc.Accept(l)
//}


func main() {
	var agent *fw.Agent
	serve := func(ws *websocket.Conn) {
		// only ONE agent server allowed to connect
		log.Debugf("new comes, agent=%+v", agent)
		if agent == nil {
			agent = fw.NewAgent(game.NewModel(), fw.NewWsReadWriter(ws), 1)
			if err := agent.Serve(); err != nil {
				log.Error(err.Error())
			}
		} else {
			log.Warning("already connected by agentsrv, cannont serve more")
		}
		// release
		agent = nil
	}

	http.Handle("/", websocket.Handler(serve))
	log.Debug("game server start on:", cfg.GameAddr())
	http.ListenAndServe(cfg.GameAddr(), nil)

}