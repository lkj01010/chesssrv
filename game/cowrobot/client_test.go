package cowrobot
import (
	"testing"
	"chess/cfg"
	"time"
	"chess/fw"
)

func TestClient1(t *testing.T){
	cli, err := fw.NewClient(cfg.AgentAddr(), &clientmodel{})
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(cli)
	go cli.Loop()

	time.Sleep(10 * time.Second)
	t.Log("end")
}
