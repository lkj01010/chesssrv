package agent
import (
	"chess/fw"
	"strconv"
	log "github.com/lkj01010/log"
	"net/rpc"
	"chess/cfg"
)

var server *Server

type Server struct {
	dao    *rpc.Client
	agents map[string]*fw.Agent
}

func NewServer() (*Server, error) {
	cli, err := rpc.Dial("tcp", cfg.DaoAddr())
	if err != nil {
		log.Error("dao client create err=", err.Error())
		return nil, err
	}
	return &Server{
		dao: cli,
		agents: make(map[string]*fw.Agent, 0),
	}, nil
}

func GetServer() *Server {
	if server != nil {
		return server
	} else {
		var err error
		server, err = NewServer()
		if err != nil {
			log.Panic("new server error: ", err.Error())
		}
	}
	return server
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
	return len(s.agents)
}

func (s *Server)Serve(rw fw.ReadWriteCloser) (err error) {
	//todo: get id
	id := strconv.Itoa(fw.FastRand())

	agent := fw.NewAgent(&handler{dao: s.dao}, rw)
	defer agent.Close()    // close it!

	s.agents[id] = agent
	if err = agent.Serve(); err != nil {
		return
	}

	delete(s.agents, id)
	return
}

func (s *Server)RemoveAgent(id string) {
	log.Fatal("jjjx")


	delete(s.agents, id)
}
