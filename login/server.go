package login
import (
	"chess/dao"
	"chess/fw"
	"strconv"
)

type server struct {
	daocli *dao.Client
	agents map[string]*fw.Agent
}

func NewServer() (*server, error) {
	cli, err := dao.NewClient()
	if err != nil {
		return nil, err
	}
	return &server{cli}, nil
}

func (s *server)serve(rw fw.ReadWriter) (err error) {
	//todo: get id
	id := strconv.Itoa(fw.FastRand())

	var agent *fw.Agent
	agent, err = fw.NewAgent(&handler{s.daocli}, fw.NewWsReadWriter(rw))
	if err != nil {
		return
	}
	s.agents[id] = agent
	if err = agent.Serve(); err != nil {
		return
	}
	return
}


