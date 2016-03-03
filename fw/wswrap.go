package fw
import (
	"golang.org/x/net/websocket"
	"github.com/lkj01010/log"
)

type WsClient struct {
	ws *websocket.Conn
}

func NewWsReadWriter(ws *websocket.Conn) *WsClient {
	return &WsClient{ws}
}

func (w *WsClient)Read(msg *string) (err error) {
	err = websocket.Message.Receive(w.ws, msg)
	log.Debugf("recv: %#v\n", *msg)
	if err != nil {
		log.Error(err)
	}
	return
}

func (w *WsClient)Write(msg string) (err error) {
	err = websocket.Message.Send(w.ws, msg)
	log.Debugf("send: %#v\n", msg)
	if err != nil {
		log.Error(err)
	}
	return
}

func (w *WsClient)Close() (err error) {
	err = w.ws.Close()
	return
}
