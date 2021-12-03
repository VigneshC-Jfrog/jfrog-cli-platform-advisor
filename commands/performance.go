package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"os"

	"github.com/jfrog/jfrog-cli-core/v2/common/commands"
	"github.com/jfrog/jfrog-cli-platform-advisor/common"
	"github.com/jfrog/jfrog-cli-platform-advisor/model"
	"github.com/wcharczuk/go-chart/v2"
)

type StorageSummary struct {
	Repo []model.Repo `json:"repositoriesSummaryList"`
}

func performanceAdvisory() error {
	getStorageSummaryAdvise()
	return nil

}

func getStorageSummaryAdvise() error {
	serverDetails, err := commands.NewCurlCommand().GetServerDetails()
	if err != nil {
		return err
	} else {
		var storageSummary StorageSummary
		var My_map = make(map[string]float64)
		var values []chart.Value
		var url = serverDetails.GetArtifactoryUrl() + "api/storageinfo"
		httpRequest := &common.HttpRequest{ReqUrl: url, ReqType: "GET", AuthUser: serverDetails.GetUser(), AuthPass: serverDetails.GetPassword()}
		var response = common.MakeHTTPCall(*httpRequest)
		if string(response) == "" {
			return errors.New("No response error")
		}
		err := json.Unmarshal(response, &storageSummary)
		if err != nil {
			errors.New("Error unmarshalling repos")
		}
		for i := 0; i < len(storageSummary.Repo); i++ {
			if storageSummary.Repo[i].RepoKey == "TOTAL" {
				continue
			}
			var size_type = strings.Split(storageSummary.Repo[i].UsedSpace, " ")[1]
			var used_size, _ = strconv.ParseFloat(strings.Split(storageSummary.Repo[i].UsedSpace, " ")[0], 64)
			if (size_type == "GB" && used_size > 0.9) || size_type == "TB" {
				fmt.Println("Reposiroty " + storageSummary.Repo[i].RepoKey + " has " + storageSummary.Repo[i].UsedSpace + ". Please split the content to different repositories.")
				My_map[storageSummary.Repo[i].RepoKey+"("+size_type+")"] = used_size

				fmt.Println("My_map: ", My_map)
			}
		}
		for l, v := range My_map {
			values = append(values, chart.Value{Label: l, Value: float64(v)})
		}
		graph := chart.BarChart{
			Title: "Test Bar Chart",
			Background: chart.Style{
				Padding: chart.Box{
					Top: 40,
				},
			},
			Height:   512,
			BarWidth: 60,
			Bars:     values,
		}

		f, _ := os.Create("output.png")
		defer f.Close()
		graph.Render(chart.PNG, f)
	}
	return nil
}
