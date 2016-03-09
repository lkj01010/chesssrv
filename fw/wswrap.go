package fw
import (
	"golang.org/x/net/websocket"
	"github.com/lkj01010/log"
)

type ReadWriteCloser interface {
	Read(msg *string) error
	Write(msg string) error
	Close() error
}

type WsClient struct {
	ws *websocket.Conn
}

func NewWsReadWriter(ws *websocket.Conn) *WsClient {
	return &WsClient{ws}
}

func (w *WsClient)Read(msg *string) (err error) {
	err = websocket.Message.Receive(w.ws, msg)
	if err != nil {
		log.Errorf("WS-RECV ERROR: %+v", err.Error())
	}else{
		log.Debugf("WS-RECV: %#v\n", *msg)
	}
	return
}

func (w *WsClient)Write(msg string) (err error) {
	err = websocket.Message.Send(w.ws, msg)
	if err != nil {
		log.Errorf("WS-SEND ERROR: %+v", err.Error())
	}else{
		log.Debugf("WS-SEND: %#v\n", msg)
	}
	return
}

func (w *WsClient)Close() (err error) {
	err = w.ws.Close()
	return
}
