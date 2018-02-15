package passport

type Provider struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
}

type Strategy struct {
	_Authenticate func(string, string) (Profile, error)
	_GetAuthURL   func(...string) (string, error)
}

type Profile map[string]interface{}

func (s *Strategy) GetAuthURL(states ...string) (string, error) {
	return s._GetAuthURL(states...)
}

func (s *Strategy) Authenticate(code, state string) (Profile, error) {
	return s._Authenticate(code, state)
}
