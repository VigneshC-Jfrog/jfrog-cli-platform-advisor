package advisories

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/jfrog/jfrog-cli-core/v2/utils/config"
	"github.com/jfrog/jfrog-cli-platform-advisor/common"
	"github.com/jfrog/jfrog-cli-platform-advisor/model"
)

type HighVirtualRepositoryAssociation struct{}

func (virutalRepoAssociation HighVirtualRepositoryAssociation) AdvisoryInfo() model.AdvisoryInfo {
	return model.AdvisoryInfo{AdvisoryName: "HighVirtualRepositoryAssociation", AdvisoryType: "performance", Severity: 2}
}

func (virutalRepoAssociation HighVirtualRepositoryAssociation) Condition(serverDetails *config.ServerDetails) {
	var repos []model.Repo
	url := serverDetails.GetArtifactoryUrl() + "api/repositories?type=virtual"
	httpRequest := &common.HttpRequest{ReqUrl: url, ReqType: "GET", AuthUser: serverDetails.GetUser(), AuthPass: serverDetails.GetPassword()}
	response := common.MakeHTTPCall(*httpRequest)
	json.Unmarshal(response, &repos)
	for i := 0; i < len(repos); i++ {
		var url = serverDetails.GetArtifactoryUrl() + "api/repositories/" + repos[i].Key
		var status string
		httpRequest := &common.HttpRequest{ReqUrl: url, ReqType: "GET", AuthUser: serverDetails.GetUser(), AuthPass: serverDetails.GetPassword()}
		var response = common.MakeHTTPCall(*httpRequest)
		var repo model.Repo
		json.Unmarshal(response, &repo)
		if len(repo.Repositories) > 2 {
			status = string("• " + strconv.Itoa(len(repo.Repositories)) + " repos [" +
				strings.Join(repo.Repositories, ",") + "] are mapped to virtual repository " + repo.Key + " Consider splitting it.\n")
			common.AddConsoleMsg(status, "warn")
		} else if len(repo.Repositories) == 0 {
			status = string("• No repositories mapped to virtual repository " + repo.Key + " Consider adding one or delete this virtual repository\n")
			common.AddConsoleMsg(status, "fatal")
		}
	}
}
