package fw
import (
	"github.com/nats-io/nats"
)

type NatsNode struct {
	conn    *nats.Conn
	SubFunc func(*nats.Msg)
}

func NewNatsNode(cfg *NatsConfig) *NatsNode {
	natsConn, err := nats.Connect(cfg.Url);
	if err != nil {
		panic("nats server connect err=" + err.Error())
	}

	for i := 0; i < len(cfg.Subs); i++ {

	}

	return NatsNode{
		conn: natsConn,
	}
}


type NatsConfig struct {
	Url  string
	Subs []string
	Pubs []string
}

