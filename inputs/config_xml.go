package inputs

import (
	"encoding/xml"
	"fmt"

	"github.com/jfrog/jfrog-cli-core/v2/common/commands"
	"github.com/jfrog/jfrog-cli-platform-advisor/common"
)

var config_xml = new(Config)

type Config struct {
	XMLName  xml.Name `xml:"config"`
	Security Security `xml:"security"`
}

type Security struct {
	AnonAccess string `xml:"anonAccessEnabled"`
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

func GetConfig() Config {
	fmt.Println("Getting config xml")
	if config_xml.Security.AnonAccess == "" {
		fmt.Println("No config xml data found retrieving it")
		data, _ := getConfigXml()
		config := &Config{}

		_ = xml.Unmarshal(data, &config)
		config_xml = config
	}
	return *config_xml
}
