package login

type LoginServer struct {
}

func (s *LoginServer)Handle(req string) (resp string, err error) {


	return "loginServer recive: " + req, nil
}
