package advisories

import (
	"encoding/json"
	"log"

	"github.com/jfrog/jfrog-cli-core/v2/utils/config"
	"github.com/jfrog/jfrog-cli-platform-advisor/common"
	"github.com/jfrog/jfrog-cli-platform-advisor/model"
)

type IsEmailAssociatedToAllUsers struct{}

func (isEmailAssociatedToAllUsers IsEmailAssociatedToAllUsers) AdvisoryInfo() model.AdvisoryInfo {
	return model.AdvisoryInfo{AdvisoryName: "IsPasswordExpirationEnabled", AdvisoryType: "security", Severity: 1}
}

func (isEmailAssociatedToAllUsers IsEmailAssociatedToAllUsers) Condition(serverDetails *config.ServerDetails) {
	url := serverDetails.GetArtifactoryUrl() + "api/security/users"
	var users []model.User
	httpRequest := &common.HttpRequest{ReqUrl: url, ReqType: "GET", AuthUser: serverDetails.GetUser(), AuthPass: serverDetails.GetPassword()}
	response := common.MakeHTTPCall(*httpRequest)
	if err := json.Unmarshal(response, &users); err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		var userInfo model.User
		url = serverDetails.GetArtifactoryUrl() + "api/security/users/" + user.Name
		httpRequest = &common.HttpRequest{ReqUrl: url, ReqType: "GET", AuthUser: serverDetails.GetUser(), AuthPass: serverDetails.GetPassword()}
		response = common.MakeHTTPCall(*httpRequest)
		if err := json.Unmarshal(response, &userInfo); err != nil {
			log.Fatal(err)
		} else {
			if userInfo.Email == "" && userInfo.Realm == "internal" && userInfo.Admin {
				common.AddConsoleMsg("Admin user ["+userInfo.Name+"] does not have a email address associated. Please set an email", "fatal")
			}
		}
	}

}
