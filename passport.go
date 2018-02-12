package passport

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

var strategyList map[string]Strategy

type Strategy struct {
	Name                   string
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

func Use(s Strategy) {
	if strategyList != nil {
		strategyList[s.Name] = s
	} else {
		strategyList = map[string]Strategy{
			s.Name: s,
		}
	}
}

func GetAuthURL(strategyName string, state string) (string, error) {
	if _, ok := strategyList[strategyName]; ok == false {
		return "", errors.New("Strategy not Registered")
	}
	strategy := strategyList[strategyName]
	paramArray := []string{}
	strategy.AuthURLParam["state"] = state
	for k, v := range strategy.AuthURLParam {
		paramArray = append(paramArray, k+"="+v)
	}
	return strategy.AuthRootURL + "?" + strings.Join(paramArray, "&"), nil
}

func Authenticate(strategyName, code, state string) (Profile, error) {
	if _, ok := strategyList[strategyName]; ok == false {
		return nil, errors.New("Strategy not Registered")
	}
	strategy := strategyList[strategyName]
	data := map[string]string{
		"client_id":     strategy.ClientID,
		"client_secret": strategy.ClientSecret,
		"code":          code,
		"redirect_uri":  strategy.RedirectURI,
		"state":         state,
		"grant_type":    "authorization_code",
	}

	bs, err := postBody(strategy.AccessTokenContentType, data, strategy.AccessTokenURL)
	if err != nil {
		fmt.Println(err)
		fmt.Println("should retry")
	}
	str := string(bs)
	accessToken := strings.Split(strings.Split(str, "&")[0], "=")[1]
	fmt.Println("accesstoken", accessToken)
	bs, err = getHttp(strategy.ProfileURL + "?access_token=" + accessToken)
	var userData map[string]interface{}
	err = json.Unmarshal(bs, &userData)
	if err != nil {
		fmt.Println(err)
	}
	return userData, nil
}
