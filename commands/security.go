package commands

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jfrog/jfrog-cli-platform-advisor/inputs"
)

type HttpRequest struct {
	ReqType  string
	ReqUrl   string
	AuthUser string
	AuthPass string
}

type Repo struct {
	Key          string   `json:"key"`
	Repositories []string `json:"repositories"`
}

func MakeHTTPCall(httpRequest HttpRequest) []byte {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", httpRequest.ReqUrl, nil)
	req.Header.Add("Authorization", "Basic "+basicAuth(httpRequest.AuthUser, httpRequest.AuthPass))
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return []byte(bodyBytes)
	}
	return nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func securityAdvisory() error {
	var output = inputs.GetConfig()
	fmt.Println(output)
	return nil
}
