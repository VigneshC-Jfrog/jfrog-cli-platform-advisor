package helper

import (
	"encoding/xml"
	"fmt"

	"github.com/jfrog/jfrog-cli-core/v2/common/commands"
	"github.com/jfrog/jfrog-cli-platform-advisor/common"
	"github.com/jfrog/jfrog-cli-platform-advisor/model"
)

var config_xml = new(model.Config)

func GetConfig() model.Config {
	fmt.Println("Getting config xml")
	if config_xml.Security.AnonAccess == "" {
		fmt.Println("No config xml data found retrieving it")
		data, _ := getConfigXml()
		config := &model.Config{}

		_ = xml.Unmarshal(data, &config)
		config_xml = config
	}
	return *config_xml
}

func getConfigXml() ([]byte, error) {
	serverDetails, err := commands.NewCurlCommand().GetServerDetails()
	if err != nil {
		return nil, err
	} else {
		var url = serverDetails.GetArtifactoryUrl() + "api/system/configuration"
		httpRequest := &common.HttpRequest{ReqUrl: url, ReqType: "GET", AuthUser: serverDetails.GetUser(), AuthPass: serverDetails.GetPassword()}
		return common.MakeHTTPCall(*httpRequest), nil
	}
}