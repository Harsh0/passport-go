package passport

func GithubStrategy(params map[string]string) Strategy {
	if params["state"] == "" {
		//generate random string
		params["state"] = "random"
	}
	return Strategy{
		Name:                   "github",
		AccessTokenURL:         "https://github.com/login/oauth/access_token",
		AccessTokenContentType: "application/json",
		ProfileURL:             "https://api.github.com/user",
		ClientID:               params["clientID"],
		ClientSecret:           params["clientSecret"],
		RedirectURI:            params["callbackURL"],
		AuthRootURL:            "https://github.com/login/oauth/authorize",
		AuthURLParam: map[string]string{
			"client_id":    params["clientID"],
			"redirect_uri": params["callbackURL"],
			"allow_signup": "true",
			"state":        params["state"],
			"scope":        "email",
		},
	}
}
