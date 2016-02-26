package dao
import (
	"net/rpc"
	"chess/cfg"
	log "github.com/lkj01010/log"
)

type Client struct {
	cli *rpc.Client
}

func NewClient() (c *Client, err error) {
	var cli *rpc.Client
	addr := cfg.IPs[cfg.RemoteType] + ":" + cfg.DaoPort
	log.Debug("addr=", addr)
	if cli, err = rpc.Dial("tcp", addr); err != nil {
		return
	}
	return &Client{cli}, nil
}

func (c *Client)Close() error {
	if err := c.cli.Close(); err != nil {
		return err
	}
	return nil
}

