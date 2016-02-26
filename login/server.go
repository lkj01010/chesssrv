package login
import (
	"chess/dao"
	"chess/fw"
	"strconv"
	log "github.com/lkj01010/log"
)

type Server struct {
	dc     *dao.Client
	agents map[string]*fw.Agent
}

func NewServer() *Server {
	dc, err := dao.NewClient()
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	return &Server{
		dc: dc,
		agents: make(map[string]*fw.Agent, 0),
	}
}

func (s *Server)Close() {
	if err := s.dc.Close(); err != nil {
		log.Error(err.Error())
	}
}

func (s *Server)AgentCount() int {
	return len(s.agents)
}

func (s *Server)Serve(rw fw.ReadWriter) (err error) {
	//todo: get id
	id := strconv.Itoa(fw.FastRand())

	var agent *fw.Agent
	agent, err = fw.NewAgent(&handler{s.dc}, rw)
	if err != nil {
		return
	}
	s.agents[id] = agent
	if err = agent.Serve(); err != nil {
		return
	}
	return
}
