package fw

import (
	"net"
	"chess/cfg"
	"golang.org/x/net/websocket"
	"github.com/lkj01010/log"
	"time"
	"fmt"
)

type ClientModel interface {
	Step(output *string)
	Handle(req string) (err error)
}

type Client struct {
	ReadWriteCloser
	ClientModel
	cmd chan string
}

func newConfig(path string) *websocket.Config {
	config, _ := websocket.NewConfig(fmt.Sprintf("ws://%s", path), "http://localhost")
	return config
}

func NewClient(addr string, m ClientModel) (*Client, error) {
	L:    client, err := net.Dial("tcp", cfg.AgentAddr())
	if err != nil {
		log.Warning("not connected to agent server, try again ...")
		time.Sleep(1 * time.Second)
		goto L
	}
	conn, err := websocket.NewClient(newConfig("/"), client)
	if err != nil {
		log.Errorf("WebSocket handshake error: %v", err)
		return nil, err
	}
	rwc := NewWsReadWriter(conn)
	return &Client{
		ReadWriteCloser: rwc,
		ClientModel: m,
		cmd: make(chan string, 5),
	}, nil
}

//func (c *Client)SendMsg(msg string) (err error) {
//	if _, err = c.Write(msg); err != nil {
//		log.Error(err.Error())
//		return
//	}
//	var rec string
//	if err = c.Read(&rec); err != nil {
//		log.Error(err.Error())
//		return
//	}
//	return
//}

func (c *Client)Loop() (err error) {
	session := make(chan string, 1)
	go func(ch chan string) {
		var buf string
		for {
			err = c.Read(&buf)
			if err != nil {
				log.Debug("client read: ", err.Error())
				return
			}
			ch <- buf
		}
	}(session)

	for {
		var output string
		c.Step(&output)
		c.Write(output)

		select {
		case cmd := <-c.cmd:
			if cmd == CtrlRemoveAgent {
				return
			}
		case msg := <-session:
			err = c.Handle(msg)
			if err != nil {
				log.Error("agent session: ", err.Error())
				return
			}
		}
	}
}

func (c *Client)Cmd(cmd string) {
	c.cmd <- cmd
}

//func (c *Client)Step() {
//	var msg string
//	msg = (`{"cmd":104,"content":"{\"account\":\"testUtf\",\"psw\":\"pswlk22\"}"}`)
//	log.Debugf("Step: send=%+v", msg)
//	c.Write(msg)
//
//	time.Sleep(3)
//}
//
//func (c *Client)Handle(req string) (resp string, err error) {
//	return "", nil
//}

