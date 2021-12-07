package advisories

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/jfrog/jfrog-cli-core/v2/utils/config"
	"github.com/jfrog/jfrog-cli-platform-advisor/common"
	"github.com/jfrog/jfrog-cli-platform-advisor/model"
)

type HighStorageSummaryPerRepo struct{}

func (highStorageSummaryPerRpo HighStorageSummaryPerRepo) AdvisoryInfo() model.AdvisoryInfo {
	return model.AdvisoryInfo{AdvisoryName: "HighStorageSummaryPerRepo", AdvisoryType: "performance", Severity: 2}
}

func (highStorageSummaryPerRepo HighStorageSummaryPerRepo) Condition(serverDetails *config.ServerDetails) {
	var storageSummary model.StorageSummary
	var storage_chart_map = make(map[string]float64)
	var url = serverDetails.GetArtifactoryUrl() + "api/storageinfo"
	httpRequest := &common.HttpRequest{ReqUrl: url, ReqType: "GET", AuthUser: serverDetails.GetUser(), AuthPass: serverDetails.GetPassword()}
	var response = common.MakeHTTPCall(*httpRequest)
	json.Unmarshal(response, &storageSummary)
	for i := 0; i < len(storageSummary.Repo); i++ {
		if storageSummary.Repo[i].RepoKey == "TOTAL" {
			continue
		}
		var size_type = strings.Split(storageSummary.Repo[i].UsedSpace, " ")[1]
		var used_size, _ = strconv.ParseFloat(strings.Split(storageSummary.Repo[i].UsedSpace, " ")[0], 64)
		if (size_type == "GB" && used_size > 1) || size_type == "TB" {
			status := "• Repository " + storageSummary.Repo[i].RepoKey + " has " + storageSummary.Repo[i].UsedSpace + ". Please split the content to different repositories.\n"
			common.AddConsoleMsg(status, "warn")
			storage_chart_map[storageSummary.Repo[i].RepoKey+"("+size_type+")"] = used_size
		}
	}
	common.GetSummaryView(storage_chart_map, "High Storage Summary Per Repo", "summary_report.png")
}
