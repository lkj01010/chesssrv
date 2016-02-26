package dao

type Server struct {
	user *User
	game *Game
}

func NewServer() *Server {
	return &Server{
		user: NewUser(),
		game: NewGame(),
	}
}

func (s *Server)Exit(){
	s.user.exit()
	s.game.exit()
}
