package cfg

const (
	Cfg = "cfg"
	Pf = "pf"
	Game = "game"
	Bill = "bill"
)

var RedisDBs = map[string]int{
	Cfg : 0,
	Pf: 1,
	Game: 2,
	Bill: 3,
}

var SrvType = "dev"

var IPs = map[string]string{
	"pro": "42.62.101.24",
	"dev": "42.62.101.24",
}

const (
	RedisPort = "6379"
	AgentPort = "13001"
	LobbyPort = "13002"
	GamePort = "13003"
	DaoPort = "13004"
)

func RedisAddr() (string) {
	addr := IPs[SrvType] + ":" + RedisPort
	return addr
}