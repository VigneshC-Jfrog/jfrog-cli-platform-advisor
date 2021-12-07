package advisories

import (
	"github.com/jfrog/jfrog-cli-core/v2/utils/config"
	"github.com/jfrog/jfrog-cli-platform-advisor/common"
	"github.com/jfrog/jfrog-cli-platform-advisor/helper"
	"github.com/jfrog/jfrog-cli-platform-advisor/model"
)

type IsAnonAccessEnabled struct{}

func (isAnonAccessEnabled IsAnonAccessEnabled) AdvisoryInfo() model.AdvisoryInfo {
	return model.AdvisoryInfo{AdvisoryName: "IsAnonAccessEnabled", AdvisoryType: "security", Severity: 1}
}

func (isAnonAccessEnabled IsAnonAccessEnabled) Condition(serverDetails *config.ServerDetails) {
	config_xml := helper.GetConfig(serverDetails)
	if config_xml.Security.AnonAccess == "true" {
		var message string = ("• Anonymous Access Enabled. Please turn off if the instance is public and not intended to have artifacts with public access")
		common.AddConsoleMsg(message, "fatal")
	} else {
		var message string = ("• Anonymous Access Disabled. We are Safe ! ")
		common.AddConsoleMsg(message, "success")
	}
}
