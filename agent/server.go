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
	dao         *rpc.Client
	game        *rpc.Client

	mu          sync.RWMutex

	//	allAgents	map[agent]interface{}
	loginAgents map[string]*fw.Agent
}

func NewServer() (*Server, error) {
	// connect dao server
	dao:    daocli, err := rpc.Dial("tcp", cfg.DaoAddr())
	if err != nil {
		log.Warningf("dao server connect fail(err=%+v), try again...", err.Error())
		time.Sleep(1 * time.Second)
		goto dao
	}

	// connect game server
	game:	gamecli, err := rpc.Dial("tcp", cfg.GameAddr())
	if err != nil {
		log.Warningf("game server connect failed(err=%+v), try again...", err.Error())
		goto game
	}

	// new server
	serverInst = &Server{
		dao: daocli,
		game: gamecli,
		loginAgents: make(map[string]*fw.Agent, 100),
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

func (s *Server)AddAgent(id string, agent *fw.Agent) {
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
	agent := fw.NewAgent(&model{dao: s.dao}, rwc)
	//	agent := NewAgent(rwc, s.dao)
	defer agent.Close()    // close it!

	if err = agent.Serve(); err != nil {
		return
	}
	return
}