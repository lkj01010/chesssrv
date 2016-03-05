package game
import (
	"net/rpc"
)

type Server struct {
	Dao    *rpc.Client

}