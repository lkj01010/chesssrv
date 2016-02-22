package login
import "encoding/json"

type LoginServer struct {
}

func (s *LoginServer)Handle(req string) (resp string, err error) {
	err := json.Unmarshal(req, )
	return "loginServer recive: " + req, nil
}
