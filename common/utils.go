package common

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/wcharczuk/go-chart/v2"
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

func GetColor(condition bool, message string) {

	if condition == true {
		red := color.New(color.FgGreen)
		red.Println(message)
	} else {
		red := color.New(color.FgRed)
		red.Println(message)
	}

}

func GetSummaryView(summarymap map[string]float64) {
	const WarningColor = "\033[1;33m%s\033[0m"
	var values []chart.Value
	for l, v := range summarymap {
		values = append(values, chart.Value{Label: l, Value: float64(v)})
	}
	graph := chart.BarChart{
		Title: "High Storage Summary Per Repo Report",
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   512,
		BarWidth: 60,
		Bars:     values,
	}
	fmt.Printf(WarningColor, "High Storage Summary Per Repo Report.....")
	fmt.Println("")
	f, _ := os.Create("summary_report.png")
	fmt.Printf(WarningColor, "High Storage Summary Per Repo Report created === summary_report.png ")
	fmt.Println("")
	defer f.Close()
	graph.Render(chart.PNG, f)
}
