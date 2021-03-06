package passport

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

/* An utility function which returns a response body.*/
func postBody(contentType string, data map[string]string, url string) ([]byte, error) {
	var strData string
	if contentType == "application/json" {
		byteData, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		strData = string(byteData)
	} else if contentType == "application/x-www-form-urlencoded" {
		strData = ConvertData(data)
	}
	payload := strings.NewReader(strData)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", contentType)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return body, nil
}

func getHttp(url string) ([]byte, error) {
	payload := strings.NewReader(string(""))
	req, _ := http.NewRequest("GET", url, payload)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return body, nil
}

func ConvertData(data map[string]string) string {
	str := ""
	for key, value := range data {
		if str != "" {
			str += "&"
		}
		str += key + "=" + value
	}
	return str
}

// GenerateRandomString will generate random stringin base64 format.
func GenerateRandomString(s int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), err
}
