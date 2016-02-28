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

const (
	rtLocal = "local"
	rtDev = "dev"
)
var RemoteType = rtLocal

var IPs = map[string]string{
	rtDev : "42.62.101.24",
	rtLocal:"127.0.0.1",
}

var RedisIPs = map[string]string{
	rtDev : "42.62.101.24",
	rtLocal: "42.62.101.24",
}

const (
	RedisPort = "6379"
	LoginPort = "13000"
	AgentPort = "13001"
	LobbyPort = "13002"
	GamePort = "13003"
	DaoPort = "13004"
)

func RedisAddr() (string) {
	addr := RedisIPs[RemoteType] + ":" + RedisPort
	return addr
}

func LoginAddr() string {
	return IPs[RemoteType] + ":" + LoginPort
}

func DaoAddr() string {
	return IPs[RemoteType] + ":" + DaoPort
}