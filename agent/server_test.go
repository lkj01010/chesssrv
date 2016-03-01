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
	client, err := net.Dial("tcp", cfg.LoginAddr())
	if err != nil {
		return nil, err
	}
	conn, err := websocket.NewClient(newConfig("/login"), client)
	if err != nil {
		log.Errorf("WebSocket handshake error: %v", err)
		return nil, err
	}
	return conn, nil
}

func startServer1() {
	serve := func(ws *websocket.Conn) {
		fmt.Printf("agent come")
		agent := fw.NewAgent(&handler{}, fw.NewWsReadWriter(ws))
		agent.Serve()
	}

	http.Handle("/login", websocket.Handler(serve))
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
		if err := server.Serve(fw.NewWsReadWriter(ws)); err != nil {
			log.Error(err.Error())
		}
		log.Infof("new agent comes, agent count=%v", server.AgentCount())
	}

	http.Handle("/", websocket.Handler(serve))
	http.ListenAndServe(cfg.LoginAddr(), nil)
	log.Info("server2 serve on ", cfg.LoginAddr())
}

var (
	serverAddr string
	once sync.Once
)

func newConfig(path string) *websocket.Config {
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
	conn, err := websocket.NewClient(newConfig("/login"), client)
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

	time.Sleep(2 * time.Second)

	conn, err := newClient()
	if err != nil {
		log.Error("newClient: ", err.Error())
		return
	}

	msg := []byte(`{"cmd":100,
		"content":"{\"account\":\"testUtf2\",\"psw\":\"pswlk在咋子。22\"}"
		}`)

//	utf8msg, _, _ := transform.String(simplifiedchinese.GBK.NewEncoder(), string(msg))
	if _, err = conn.Write(msg); err != nil {
		log.Error(err.Error())
		return
	}
	var rec string
	if err = websocket.Message.Receive(conn, &rec); err != nil {
		log.Error(err.Error())
		return
	}

}