package advisories

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/jfrog/jfrog-cli-core/v2/common/commands"
	"github.com/jfrog/jfrog-cli-platform-advisor/common"
	"github.com/jfrog/jfrog-cli-platform-advisor/model"
)

type HighVirualRepositoryAssociation struct{}

func (virutalRepoAssociation HighVirualRepositoryAssociation) AdvisoryInfo() model.AdvisoryInfo {
	return model.AdvisoryInfo{AdvisoryName: "HighVirualRepositoryAssociation", AdvisoryType: "performance", Severity: 2}
}

func (virutalRepoAssociation HighVirualRepositoryAssociation) Condition() bool {
	serverDetails, err := commands.NewCurlCommand().GetServerDetails()
	if err != nil {
		return false
	}
	var result bool = true
	var repos []model.Repo
	url := serverDetails.GetArtifactoryUrl() + "api/repositories?type=virtual"
	httpRequest := &common.HttpRequest{ReqUrl: url, ReqType: "GET", AuthUser: serverDetails.GetUser(), AuthPass: serverDetails.GetPassword()}
	response := common.MakeHTTPCall(*httpRequest)
	json.Unmarshal(response, &repos)
	for i := 0; i < len(repos); i++ {
		var url = serverDetails.GetArtifactoryUrl() + "api/repositories/" + repos[i].Key
		httpRequest := &common.HttpRequest{ReqUrl: url, ReqType: "GET", AuthUser: serverDetails.GetUser(), AuthPass: serverDetails.GetPassword()}
		var response = common.MakeHTTPCall(*httpRequest)
		var repo model.Repo
		json.Unmarshal(response, &repo)
		if len(repo.Repositories) > 4 {
			result = false
			fmt.Println(strconv.Itoa(len(repo.Repositories)) + " are mapped to virtual repository " + repo.Key + " Consider splitting it.")
		} else if len(repo.Repositories) == 0 {
			fmt.Println("No repositories mapped to virtual repository " + repo.Key + " Consider adding one or delete this virtual repository")
		}
	}
	return result
}
