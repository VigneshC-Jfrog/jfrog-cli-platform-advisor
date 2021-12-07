package common

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gookit/color"
	"github.com/wcharczuk/go-chart/v2"
)

var colour_map = map[string]color.Color{"info": color.FgLightBlue, "warn": color.FgYellow, "fatal": color.FgRed, "success": color.FgGreen, "default": color.White}

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
	} else {
		log.Fatal("Error making REST call to url " + httpRequest.ReqUrl + " with response " + resp.Status)
	}
	return nil
}

func GetSummaryView(summarymap map[string]float64, title, file_name string) {
	var values []chart.Value
	for l, v := range summarymap {
		values = append(values, chart.Value{Label: l, Value: float64(v)})
	}
	graph := chart.BarChart{
		Title: title,
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   512,
		BarWidth: 200,
		Bars:     values,
	}
	f, _ := os.Create(file_name)
	defer f.Close()
	graph.Render(chart.PNG, f)
	AddConsoleMsg("File created: "+file_name+"\n", "info")
}

func AddConsoleMsg(message string, level string) {
	decor_print := color.New(colour_map[level])
	decor_print.Println(time.Now().Format("01-02-2006 15:04:05") + "\t" + message)
}
