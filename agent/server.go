package agent
import (
	"chess/fw"
	log "github.com/lkj01010/log"
	"net/rpc"
	"chess/cfg"
	"sync"
	"time"
)

var serverInst *Server

type Server struct {
	dao       *rpc.Client
	//	game        *rpc.Client
	gameCli   *fw.Client
	gameCtrl  chan string

	connIdAcc int
	mu        sync.RWMutex
	agents    map[int]*fw.Agent
}

func NewServer() (*Server, error) {
	// connect dao server
	dao:    daocli, err := rpc.Dial("tcp", cfg.DaoAddr())
	if err != nil {
		log.Warningf("dao server connect fail(err=%+v), try again...", err.Error())
		time.Sleep(1 * time.Second)
		goto dao
	}

	ctrl := make(chan string, 10)
	game:    cli, err := fw.NewClient(cfg.GameAddr(), &gameCliModel{}, ctrl)
	if err != nil {
		log.Error(err.Error())
		log.Warningf("game server connect fail(err=%+v), try again...", err.Error())
		time.Sleep(1 * time.Second)
		goto game
	}
	go cli.Loop()

	// daemon routine (to be done), with "ctrl"
	// ......
	////////////////////

	// new server
	serverInst = &Server{
		dao: daocli,
		gameCli: cli,
		gameCtrl: ctrl,
		connIdAcc: 0,
		agents: make(map[int]*fw.Agent, 1000),
	}
	return serverInst, nil
}

func (s *Server)Close() {
	if err := s.dao.Close(); err != nil {
		log.Error(err.Error())
	}
	for _, v := range s.agents {
		if err := v.Close(); err != nil {
			log.Error(err.Error())
		}
	}
}

func (s *Server)AgentCount() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.agents)
}

func (s *Server)GetAgent(connId int) *fw.Agent {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if a, ok := s.agents[connId]; ok {
		return a
	}else{
		return nil
	}
}

func (s *Server)Serve(rwc fw.ReadWriteCloser) (err error) {
	s.connIdAcc ++
	agent := fw.NewAgent(&model{dao: s.dao}, rwc, s.connIdAcc)

	// add to map
	s.mu.Lock()
	s.agents[s.connIdAcc] = agent
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		delete(s.agents, agent.ConnId)
		s.mu.Unlock()
		defer agent.Close()    // close it!
	}()

	if err = agent.Serve(); err != nil {
		return
	}
	return
}