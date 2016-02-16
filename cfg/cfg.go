package cfg

var RedisDBs = map[string]int{
	"cfg" : 0,
	"game": 1,
	"bill": 2,
}

var SrvType = "dev"

var IPs = map[string]string{
	"pro": "42.62.101.24",
	"dev": "42.62.101.24",
}

var (
	RedisPort = "6379"
	AgentPort = "13001"
	LobbyPort = "13002"
	GamePort = "13003"
	DataPort = "13004"
)