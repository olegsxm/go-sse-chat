package services

type Services struct {
	auth AuthService
}

func (s *Services) Auth() AuthService {
	return s.auth
}

func (s *Services) Chat() {
	//TODO implement me
	panic("implement me")
}

func (s *Services) Message() {
	//TODO implement me
	panic("implement me")
}

func New() Services {
	return Services{
		auth: newAuthService(),
	}
}
