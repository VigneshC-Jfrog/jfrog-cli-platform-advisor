package advisories

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/jfrog/jfrog-cli-core/v2/common/commands"
	"github.com/jfrog/jfrog-cli-platform-advisor/common"
	"github.com/jfrog/jfrog-cli-platform-advisor/model"
)

type HighStorageSummaryPerRepo struct{}

func (highStorageSummaryPerRpo HighStorageSummaryPerRepo) AdvisoryInfo() model.AdvisoryInfo {
	return model.AdvisoryInfo{AdvisoryName: "HighStorageSummaryPerRepo", AdvisoryType: "performance", Severity: 2}
}

func (highStorageSummaryPerRpo HighStorageSummaryPerRepo) Condition() bool {
	serverDetails, _ := commands.NewCurlCommand().GetServerDetails()
	var result bool = true
	var storageSummary model.StorageSummary
	var My_map = make(map[string]float64)
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
		if (size_type == "GB" && used_size > 3) || size_type == "TB" {
			result = false
			fmt.Println("Reposiroty " + storageSummary.Repo[i].RepoKey + " has " + storageSummary.Repo[i].UsedSpace + ". Please split the content to different repositories.")
			My_map[storageSummary.Repo[i].RepoKey+"("+size_type+")"] = used_size
		}
	}
	return result
}
