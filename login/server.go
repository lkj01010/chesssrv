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

func NewServer() (*Server, error) {
	dc, err := dao.NewClient()
	if err != nil {
		log.Error("dao client create err=", err.Error())
		return nil, err
	}
	return &Server{
		dc: dc,
		agents: make(map[string]*fw.Agent, 0),
	}, nil
}

func (s *Server)Close() {
	if err := s.dc.Close(); err != nil {
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

	agent := fw.NewAgent(&handler{s.dc}, rw)
	defer agent.Close()	// close it!

	s.agents[id] = agent
	if err = agent.Serve(); err != nil {
		return
	}

	delete(s.agents, id)
	return
}
