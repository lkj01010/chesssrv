package main

import (
	"net/rpc"
	"net"
	"chess/cfg"
	"chess/game"
	"github.com/lkj01010/log"
)

func main() {
	//register rpc
	m := game.NewModel()
	rpc.Register(m.Server)
	defer func() {
		m.Exit()
	}()

	//network
	serverAddr := cfg.GameAddr()
	l, e := net.Listen("tcp", serverAddr)
	if e != nil {
		log.Fatalf("net.Listen tcp : %v", e)
	}
	log.Info("dao RPC server listening on ", serverAddr)
	rpc.Accept(l)

	// todo: replace rpc to tcp
}