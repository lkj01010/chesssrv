package fw
import (
	"net/rpc"
	"github.com/lkj01010/log"
)

type RpcClient struct {
	*rpc.Client
}

func NewRpcClient(addr string) (c *RpcClient, err error) {
	var cli *rpc.Client
	if cli, err = rpc.Dial("tcp", addr); err != nil {
		return
	}
	return &RpcClient{cli}, nil
}

func (c *RpcClient)Close() error {
	log.Debug("rpc client close")
	if err := c.Close(); err != nil {
		return err
	}
	return nil
}