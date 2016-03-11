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
var RemoteType = rtDev

//const (
//	ipSpecial = "spe"
//	ipLocal = "loc"
//)
//var IpType = ipSpecial

var IPs = map[string]string{
	rtDev : "",
	rtLocal:"",
}

var RedisIPs = map[string]string{
	rtDev : "42.62.101.24",
	rtLocal: "",
}

const (
	RedisPort = "6379"
	AgentPort = "13001"
	LobbyPort = "13002"
	GamePort = "13003"
	DaoPort = "13004"
)

func RedisAddr() (string) {
	return RedisIPs[RemoteType] + ":" + RedisPort
}

func AgentAddr() string {
	return IPs[RemoteType] + ":" + AgentPort
}

func LobbyAddr() string{
	return IPs[RemoteType] + ":" + LobbyPort
}

func GameAddr() string{
	return IPs[RemoteType] + ":" + GamePort
}

func DaoAddr() string {
	return IPs[RemoteType] + ":" + DaoPort
}