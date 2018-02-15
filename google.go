package passport

func GoogleStrategy(params map[string]string) Strategy {
	if params["state"] == "" {
		//generate random string
		params["state"] = "random"
	}
	return Strategy{
		AccessTokenURL:         "https://accounts.google.com/o/oauth2/token",
		AccessTokenContentType: "application/x-www-form-urlencoded",
		ProfileURL:             "https://www.googleapis.com/plus/v1/people/me",
		ClientID:               params["clientID"],
		ClientSecret:           params["clientSecret"],
		RedirectURI:            params["callbackURL"],
		AuthRootURL:            "https://accounts.google.com/o/oauth2/auth",
		AuthURLParam: map[string]string{
			"client_id":     params["clientID"],
			"response_type": "code",
			"scope":         "email",
			"redirect_uri":  params["callbackURL"],
			"allow_signup":  "true",
			"state":         params["state"],
		},
	}
}
