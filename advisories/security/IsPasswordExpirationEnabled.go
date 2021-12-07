package advisories

import (
	"encoding/json"
	"log"

	"github.com/jfrog/jfrog-cli-core/v2/utils/config"
	"github.com/jfrog/jfrog-cli-platform-advisor/common"
	"github.com/jfrog/jfrog-cli-platform-advisor/model"
)

type IsPasswordExpirationEnabled struct{}

func (isPasswordExpirationEnabled IsPasswordExpirationEnabled) AdvisoryInfo() model.AdvisoryInfo {
	return model.AdvisoryInfo{AdvisoryName: "IsPasswordExpirationEnabled", AdvisoryType: "security", Severity: 1}
}

func (isPasswordExpirationEnabled IsPasswordExpirationEnabled) Condition(serverDetails *config.ServerDetails) {
	url := serverDetails.GetArtifactoryUrl() + "api/security/configuration/passwordExpirationPolicy"
	httpRequest := &common.HttpRequest{ReqUrl: url, ReqType: "GET", AuthUser: serverDetails.GetUser(), AuthPass: serverDetails.GetPassword()}
	response := common.MakeHTTPCall(*httpRequest)
	var objmap map[string]interface{}
	if err := json.Unmarshal(response, &objmap); err != nil {
		log.Fatal(err)
	}
	isExpiryEnabled := objmap["enabled"]
	if isExpiryEnabled == "false" {
		common.AddConsoleMsg("Password expiration policy disabled", "fatal")
	} else {
		common.AddConsoleMsg("Password expiration policy enabled", "success")
	}
}
