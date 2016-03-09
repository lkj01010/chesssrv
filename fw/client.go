package fw

import (
	"net"
	"golang.org/x/net/websocket"
	"github.com/lkj01010/log"
	"time"
	"fmt"
)

const (
	ClientReturn = "ClientReturn"
)

type ClientModel interface {
	Handle(req string) (err error)
}

type Client struct {
	ReadWriteCloser
	ClientModel
	send chan string
	ctrl chan string
}

func newConfig(path string) *websocket.Config {
	config, _ := websocket.NewConfig(fmt.Sprintf("ws://%s", path), "http://localhost")
	return config
}

func NewClient(addr string, m ClientModel, out chan string) (*Client, error) {
	L:    client, err := net.Dial("tcp", addr)
	if err != nil {
		log.Warningf("not connected to %+v, try again ...", addr)
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
		send: make(chan string, 10),
		ctrl: out,
	}, nil
}

func (c *Client)Loop() (err error) {
	defer func() {
		c.ctrl <- ClientReturn
	}()

	session := make(chan string, 1)
	go func(ch chan string) {
		var buf string
		for {
			err = c.Read(&buf)
			if err != nil {
				log.Error("client read failed: ", err.Error())
				return
			}
			ch <- buf
		}
	}(session)

	for {
		select {
		case cmd := <-c.ctrl:
			if cmd == CtrlRemoveAgent {
				return
			}
		case msg := <-session:
			err = c.Handle(msg)
			if err != nil {
				log.Error("client handle failed: ", err.Error())
				return
			}
		case msg := <-c.send:
			if err = c.Write(msg); err != nil {
				// 写错误
				log.Error("client write failed: ", err.Error())
				return
			}
		}
	}
}

func (c *Client)Cmd(cmd string) {
	c.ctrl <- cmd
}

func (c *Client)Send(msg string) {
	c.send <- msg
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

