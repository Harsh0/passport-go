package passport

import (
	"encoding/json"
	"fmt"
	"strings"
)

func GoogleStrategy(params map[string]string) Strategy {
	return Strategy{
		ClientID:     params["clientID"],
		ClientSecret: params["clientSecret"],
		RedirectURI:  params["callbackURL"],
		AuthRootURL:  "https://accounts.google.com/o/oauth2/auth",
		AuthURLParam: map[string]string{
			"client_id":     params["clientID"],
			"response_type": "code",
			"scope":         "email",
			"redirect_uri":  params["callbackURL"],
			"allow_signup":  "true",
			"state":         "",
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

			bs, err := postBody("application/x-www-form-urlencoded", data, "https://accounts.google.com/o/oauth2/token")
			if err != nil {
				fmt.Println(err)
				fmt.Println("should retry")
			}
			str := string(bs)
			accessToken := strings.Split(strings.Split(str, "&")[0], "=")[1]
			fmt.Println("accesstoken", accessToken)
			bs, err = getHttp("https://www.googleapis.com/plus/v1/people/me" + "?access_token=" + accessToken)
			var userData map[string]interface{}
			err = json.Unmarshal(bs, &userData)
			if err != nil {
				fmt.Println(err)
			}
			return userData, nil
		},
	}
}
