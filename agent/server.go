package agent
import (
	"chess/fw"
	log "github.com/lkj01010/log"
	"net/rpc"
	"chess/cfg"
	"sync"
)

var serverInst *Server

type Server struct {
	dao         *rpc.Client

	mu          sync.RWMutex

	//	allAgents	map[agent]interface{}
	//todo: 相关逻辑和操作函数
	loginAgents map[string]*agent
}

func NewServer() (*Server, error) {
	cli, err := rpc.Dial("tcp", cfg.DaoAddr())
	if err != nil {
		log.Error("dao client create err=", err.Error())
		return nil, err
	}
	serverInst = &Server{
		dao: cli,
		loginAgents: make(map[string]*agent, 0),
	}
	return serverInst, nil
}

//func GetServer() *Server {
//	if serverInst != nil {
//		return serverInst
//	} else {
//		var err error
//		serverInst, err = NewServer()
//		if err != nil {
//			log.Panic("new server error: ", err.Error())
//		}
//	}
//	return serverInst
//}

func (s *Server)Close() {
	if err := s.dao.Close(); err != nil {
		log.Error(err.Error())
	}
	for _, v := range s.loginAgents {
		if err := v.Close(); err != nil {
			log.Error(err.Error())
		}
	}
}

func (s *Server)AddAgent(id string, agent *agent) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.loginAgents[id]; ok {
		log.Warning("AddAgent: agent exist: id=", id)
	}
	s.loginAgents[id] = agent
	log.Debugf("agent add, agent count=%v", len(s.loginAgents))
}

func (s *Server)RemoveAgent(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.loginAgents[id]; ok {
		delete(s.loginAgents, id)
		log.Debugf("agent remove, agent count=%v", len(s.loginAgents))
	} else {
		log.Warning("RemoveAgent: agent not exist: id=", id)
	}
}

func (s *Server)AgentCount() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.loginAgents)
}

func (s *Server)Serve(rwc fw.ReadWriteCloser) (err error) {
	//	agent := fw.NewAgent(&model{dao: s.dao}, rw)
	agent := NewAgent(rwc, s.dao)
	defer agent.Close()    // close it!

	if err = agent.Serve(); err != nil {
		return
	}
	return
}