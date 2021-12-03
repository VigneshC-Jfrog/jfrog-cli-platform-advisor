package common

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
)

type HttpRequest struct {
	ReqType  string
	ReqUrl   string
	AuthUser string
	AuthPass string
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func MakeHTTPCall(httpRequest HttpRequest) []byte {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", httpRequest.ReqUrl, nil)
	req.Header.Add("Authorization", "Basic "+basicAuth(httpRequest.AuthUser, httpRequest.AuthPass))
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return []byte(bodyBytes)
	}
	return nil
}
