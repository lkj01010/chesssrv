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
	serve := func(ws *websocket.Conn) {
		agent := fw.NewAgent(game.NewModel(), fw.NewWsReadWriter(ws))
		if err := agent.Serve(); err != nil {
			log.Error(err.Error())
		}
	}

	http.Handle("/", websocket.Handler(serve))
	log.Debug("game server start on:", cfg.GameAddr())
	http.ListenAndServe(cfg.GameAddr(), nil)

}