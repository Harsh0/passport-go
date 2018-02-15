package passport

import (
	"strings"
)

type Strategy struct {
	AccessTokenURL         string
	AccessTokenContentType string
	ProfileURL             string
	ClientID               string
	ClientSecret           string
	RedirectURI            string
	AuthRootURL            string
	AuthURLParam           map[string]string
	_Authenticate          func(string, string) (Profile, error)
}

type Profile map[string]interface{}

func (s *Strategy) GetAuthURL(states ...string) (string, error) {
	paramArray := []string{}
	var state string
	if len(states) == 0 || states[0] == "" {
		state = "changestatehere"
	} else {
		state = states[0]
	}
	s.AuthURLParam["state"] = state
	for k, v := range s.AuthURLParam {
		paramArray = append(paramArray, k+"="+v)
	}
	return s.AuthRootURL + "?" + strings.Join(paramArray, "&"), nil
}

func (s *Strategy) Authenticate(code, state string) (Profile, error) {
	return s._Authenticate(code, state)
}
