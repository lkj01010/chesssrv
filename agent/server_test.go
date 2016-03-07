package agent
import (
	"testing"
	"golang.org/x/net/websocket"
	"fmt"
	"chess/fw"
	"net/http"
	"net"
	"bytes"
	"sync"
	"net/http/httptest"
	"github.com/lkj01010/log"
	"chess/cfg"
	"time"
)

func newClient() (*websocket.Conn, error) {
	//todo
L:	client, err := net.Dial("tcp", cfg.AgentAddr())
	if err != nil {
		log.Warning("not connected to agent server, try again ...")
		time.Sleep(1)
		goto L
	}
	conn, err := websocket.NewClient(newConfig_("/"), client)
	if err != nil {
		log.Errorf("WebSocket handshake error: %v", err)
		return nil, err
	}
	return conn, nil
}

func sendMsg(conn *websocket.Conn, msg []byte) (err error){
	if _, err = conn.Write(msg); err != nil {
		log.Error(err.Error())
		return
	}
	var rec string
	if err = websocket.Message.Receive(conn, &rec); err != nil {
		log.Error(err.Error())
		return
	}
	return
}

func startServer1() {
	serve := func(ws *websocket.Conn) {
		fmt.Printf("agent come")
		agent := fw.NewAgent(&model{}, fw.NewWsReadWriter(ws))
		agent.Serve()
	}

	http.Handle("/agent", websocket.Handler(serve))
	//	http.ListenAndServe(":8000", nil)

	server := httptest.NewServer(nil)
	serverAddr = server.Listener.Addr().String()
	log.Info("Test WebSocket server listening on ", serverAddr)
}

func startServer2() {
	server, err := NewServer()
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer func() {
		server.Close()
	}()

	serve := func(ws *websocket.Conn) {
		log.Debug("new agent")
		if err := server.Serve(fw.NewWsReadWriter(ws)); err != nil {
			log.Error(err.Error())
		}
		log.Debug("agent leave")
	}

	http.Handle("/", websocket.Handler(serve))
	http.ListenAndServe(cfg.AgentAddr(), nil)
	log.Info("server2 serve on ", cfg.AgentAddr())
}

var (
	serverAddr string
	once sync.Once
)

func newConfig_(path string) *websocket.Config {
	log.Infof("serverAddr=%+v, path=%+v", serverAddr, path)
	config, _ := websocket.NewConfig(fmt.Sprintf("ws://%s%s", serverAddr, path), "http://localhost")
	return config
}

func TestServer1(t *testing.T) {
	once.Do(startServer1)

	// websocket.Dial()
	client, err := net.Dial("tcp", serverAddr)
	if err != nil {
		t.Fatal("dialing", err)
	}
	log.Info("t=%v", t)
	conn, err := websocket.NewClient(newConfig_("/agent"), client)
	if err != nil {
		t.Errorf("WebSocket handshake error: %v", err)
		return
	}

	for i := 0; i < 10; i++ {
		msg := []byte("hello, world")
		fw.PrintType(msg, "msg")
		msg = append(msg, byte(i))
		//		append(msg, []byte("\n"))
		if _, err := conn.Write(msg); err != nil {
			t.Errorf("Write: %v", err)
		}
		var actual_msg = make([]byte, 512)
		n, err := conn.Read(actual_msg)
		if err != nil {
			t.Errorf("Read: %v", err)
		}
		actual_msg = actual_msg[0:n]
		if !bytes.Equal(msg, actual_msg) {
			t.Logf("Test: send %q got %q", msg, actual_msg)
		}
	}

	conn.Close()
}

func TestServer2(t *testing.T) {
	go func() {
		once.Do(startServer2)
	}()

//	time.Sleep(5 * time.Second)

	conn, err := newClient()
	defer conn.Close()

	if err != nil {
		log.Error("newClient: ", err.Error())
		return
	}

	var msg []byte

//	log.Info("test:register")
//	msg = []byte(`{"cmd":100,
//		"content":"{\"account\":\"testUtf\",\"psw\":\"pswlk22\"}"
//		}`)
//	sendMsg(conn, msg)
//	time.Sleep(1 * time.Second)

	log.Info("test:cmdLoginReq1")
	msg = []byte(`{"cmd":104,"content":"{\"account\":\"testUtf\",\"psw\":\"pswlk22\"}"}`)
	sendMsg(conn, msg)
	time.Sleep(1 * time.Second)

//	log.Info("test:cmdLoginReq2")
//	msg = []byte(`{"cmd":104,"content":"{\"account\":\"testUtf\",\"psw\":\"p\"}"}`)
//	sendMsg(conn, msg)
//	time.Sleep(1 * time.Second)

	log.Info("test:handleInfo")
	msg = []byte(`{"cmd":106,"content":"{\"account\":\"testUtf\",\"psw\":\"p\"}"}`)
	sendMsg(conn, msg)
	time.Sleep(11 * time.Second)

}

func TestServer3(t *testing.T) {


}

const (
	vv01 = iota
	vv02
	vv03
	vv04
)

const (
	tt01 = iota << 2
	tt02
	tt03
	tt04
	tt05
	tt06
)
func Test3(t *testing.T){
	log.Debug(tt01)
	log.Debug(tt02)
	log.Debug(tt03)
	log.Debug(tt04)
	log.Debug(tt05)
	log.Debug(tt06)


	log.Debug(tt01 >> 2)
	log.Debug(tt02 >> 2)
	log.Debug(tt03 >> 2)
	log.Debug(tt04 >> 2)
	log.Debug(tt05 >> 2)
	log.Debug(tt06 >> 2)
	log.Debug("---")

	log.Debug(vv02+tt03)
	log.Debug(vv02+tt04)
	log.Debug(vv02+tt05)
	log.Debug(vv02+tt06)
}