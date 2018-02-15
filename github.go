package passport

import (
	"encoding/json"
	"fmt"
	"strings"
)

func GithubStrategy(params map[string]string) Strategy {
	return Strategy{
		ClientID:     params["clientID"],
		ClientSecret: params["clientSecret"],
		RedirectURI:  params["callbackURL"],
		AuthRootURL:  "https://github.com/login/oauth/authorize",
		AuthURLParam: map[string]string{
			"client_id":    params["clientID"],
			"redirect_uri": params["callbackURL"],
			"allow_signup": "true",
			"state":        "",
			"scope":        "email",
		},
		_Authenticate: func(code, state string) (Profile, error) {
			data := map[string]string{
				"client_id":     params["clientID"],
				"client_secret": params["clientSecret"],
				"code":          code,
				"redirect_uri":  params["callbackURL"],
				"state":         state,
				"grant_type":    "authorization_code",
			}

			bs, err := postBody("application/json", data, "https://github.com/login/oauth/access_token")
			if err != nil {
				fmt.Println(err)
				fmt.Println("should retry")
			}
			str := string(bs)
			accessToken := strings.Split(strings.Split(str, "&")[0], "=")[1]
			fmt.Println("accesstoken", accessToken)
			bs, err = getHttp("https://api.github.com/user" + "?access_token=" + accessToken)
			var userData map[string]interface{}
			err = json.Unmarshal(bs, &userData)
			if err != nil {
				fmt.Println(err)
			}
			return userData, nil
		},
	}
}
