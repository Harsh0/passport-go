package passport

import (
	"encoding/json"
	"strings"
)

func GithubStrategy(provider Provider) Strategy {
	return Strategy{
		_GetAuthURL: func(states ...string) (string, error) {
			AuthURLParam := map[string]string{
				"client_id":    provider.ClientID,
				"redirect_uri": provider.RedirectURI,
				"allow_signup": "true",
				"state":        "",
				"scope":        "email",
			}
			paramArray := []string{}
			var state string
			if len(states) == 0 || states[0] == "" {
				state = GenerateRandomString(10)
			} else {
				state = states[0]
			}
			s.AuthURLParam["state"] = state
			for k, v := range AuthURLParam {
				paramArray = append(paramArray, k+"="+v)
			}
			return "https://github.com/login/oauth/authorize" + "?" + strings.Join(paramArray, "&"), nil
		},
		_Authenticate: func(code, state string) (Profile, error) {
			data := map[string]string{
				"client_id":     provider.ClientID,
				"client_secret": provider.ClientSecret,
				"code":          code,
				"redirect_uri":  provider.RedirectURI,
				"state":         state,
				"grant_type":    "authorization_code",
			}
			bs, err := postBody("application/json", data, "https://github.com/login/oauth/access_token")
			if err != nil {
				return nil, err
			}
			str := string(bs)
			accessToken := strings.Split(strings.Split(str, "&")[0], "=")[1]
			bs, err = getHttp("https://api.github.com/user" + "?access_token=" + accessToken)
			var userData map[string]interface{}
			err = json.Unmarshal(bs, &userData)
			if err != nil {
				return nil, err
			}
			userData["access_token"] = accessToken
			return userData, nil
		},
	}
}
