package common

import (
	"github.com/jfrog/jfrog-cli-core/v2/artifactory/commands/curl"
	coreCommonCommands "github.com/jfrog/jfrog-cli-core/v2/common/commands"
)

func newRtCurlCommand(args []string) (*curl.RtCurlCommand, error) {
	curlCommand := coreCommonCommands.NewCurlCommand().SetArguments(args)
	rtCurlCommand := curl.NewRtCurlCommand(*curlCommand)
	rtDetails, err := rtCurlCommand.GetServerDetails()
	if err != nil {
		return nil, err
	}
	rtCurlCommand.SetServerDetails(rtDetails)
	rtCurlCommand.SetUrl(rtDetails.ArtifactoryUrl)
	return rtCurlCommand, err
}

func GetConfigXml() {
	var args = []string{"-XGET", "/api/build"}
	newRtCurlCommand(args)
}
