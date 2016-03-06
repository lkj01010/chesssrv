package cowrobot
import (
//	"github.com/lkj01010/log"
//	"golang.org/x/net/websocket"
//	"time"
//	"chess/cfg"
//	"net"
)


type client struct {

}

func (c *client)connect() {

}
//func newClient() (*websocket.Conn, error) {
//	L:    client, err := net.Dial("tcp", cfg.AgentAddr())
//	if err != nil {
//		log.Warning("not connected to agent server, try again ...")
//		time.Sleep(1)
//		goto L
//	}
//	conn, err := websocket.NewClient(newConfig("/"), client)
//	if err != nil {
//		log.Errorf("WebSocket handshake error: %v", err)
//		return nil, err
//	}
//	return conn, nil
//}