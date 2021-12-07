package helper

import (
	"encoding/xml"

	"github.com/jfrog/jfrog-cli-core/v2/utils/config"
	"github.com/jfrog/jfrog-cli-platform-advisor/common"
	"github.com/jfrog/jfrog-cli-platform-advisor/model"
)

var config_xml = new(model.Config)

func GetConfig(serverDetails *config.ServerDetails) model.Config {
	if config_xml.Security.AnonAccess == "" {
		var url = serverDetails.GetArtifactoryUrl() + "api/system/configuration"
		httpRequest := &common.HttpRequest{ReqUrl: url, ReqType: "GET", AuthUser: serverDetails.GetUser(), AuthPass: serverDetails.GetPassword()}
		data := common.MakeHTTPCall(*httpRequest)
		config := &model.Config{}

		_ = xml.Unmarshal(data, &config)
		config_xml = config
	}
	return *config_xml
}
