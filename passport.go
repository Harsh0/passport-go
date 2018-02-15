package passport

import (
	"encoding/json"
	"fmt"
	"strings"
)

var strategyList map[string]Strategy

type Provider interface {
	GetAuthURL(...string) string
	Authenticate() (Profile, error)
}

type Strategy struct {
	AccessTokenURL         string
	AccessTokenContentType string
	ProfileURL             string
	ClientID               string
	ClientSecret           string
	RedirectURI            string
	AuthRootURL            string
	AuthURLParam           map[string]string
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
	data := map[string]string{
		"client_id":     s.ClientID,
		"client_secret": s.ClientSecret,
		"code":          code,
		"redirect_uri":  s.RedirectURI,
		"state":         state,
		"grant_type":    "authorization_code",
	}

	bs, err := postBody(s.AccessTokenContentType, data, s.AccessTokenURL)
	if err != nil {
		fmt.Println(err)
		fmt.Println("should retry")
	}
	str := string(bs)
	accessToken := strings.Split(strings.Split(str, "&")[0], "=")[1]
	fmt.Println("accesstoken", accessToken)
	bs, err = getHttp(s.ProfileURL + "?access_token=" + accessToken)
	var userData map[string]interface{}
	err = json.Unmarshal(bs, &userData)
	if err != nil {
		fmt.Println(err)
	}
	return userData, nil
}
