package cfg

var RedisDBs = map[string]int{
	"cfg" : 0,
	"platform": 1,
	"game": 2,
	"bill": 3,
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
	DataPort = "13004"
)

func RedisAddr() (string) {
	addr := IPs[SrvType] + ":" + RedisPort
	return addr
}