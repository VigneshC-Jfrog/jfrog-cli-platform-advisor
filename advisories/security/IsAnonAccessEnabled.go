package advisories

import (
	"github.com/jfrog/jfrog-cli-platform-advisor/common"
	"github.com/jfrog/jfrog-cli-platform-advisor/helper"
	"github.com/jfrog/jfrog-cli-platform-advisor/model"
)

type IsAnonAccessEnabled struct{}

func (isAnonAccessEnabled IsAnonAccessEnabled) AdvisoryInfo() model.AdvisoryInfo {
	return model.AdvisoryInfo{AdvisoryName: "IsAnonAccessEnabled", AdvisoryType: "security", Severity: 1}
}

func (isAnonAccessEnabled IsAnonAccessEnabled) Condition() bool {
	var result bool = true
	config_xml := helper.GetConfig()
	if config_xml.Security.AnonAccess == "true" {
		var message string = ("• Anonymous Access Enabled. Please turn off if the instance is public and not intended to have artifacts with public access")
		result = false
		common.GetColor(result, message)
		result = false
	} else {
		var message string = ("• Anonymous Access Disabled. We are Safe ! ")
		common.GetColor(result, message)
	}
	return result
}
