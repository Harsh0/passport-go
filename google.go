package passport

import (
	"encoding/json"
	"strings"
)

func GoogleStrategy(provider Provider) Strategy {
	return Strategy{
		_GetAuthURL: func(states ...string) (string, error) {
			AuthURLParam = map[string]string{
				"client_id":     provider.ClientID,
				"response_type": "code",
				"scope":         "email",
				"redirect_uri":  provider.RedirectURI,
				"allow_signup":  "true",
				"state":         "",
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
			return "https://accounts.google.com/o/oauth2/auth" + "?" + strings.Join(paramArray, "&"), nil
		},
		_Authenticate: func(code, state string) (Profile, error) {
			data := map[string]string{
				"client_id":     provider.ClientID,
				"client_secret": provider.ClientSecret,
				"code":          code,
				"redirect_uri":  provider.RedirectURI,
				"grant_type":    "authorization_code",
			}
			bs, err := postBody("application/x-www-form-urlencoded", data, "https://accounts.google.com/o/oauth2/token")
			if err != nil {
				return nil, err
			}
			var authData map[string]interface{}
			err = json.Unmarshal(bs, &authData)
			if err != nil {
				return nil, err
			}
			accessToken := authData["access_token"].(string)
			bs, err = getHttp("https://www.googleapis.com/plus/v1/people/me" + "?access_token=" + accessToken)
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
