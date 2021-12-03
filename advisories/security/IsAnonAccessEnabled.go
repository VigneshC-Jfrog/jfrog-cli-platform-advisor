package advisories

import (
	"fmt"

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
		fmt.Println("Anonymous Access Enabled. Please turn off if the instance is public and not intended to have artifacts with public access")
		result = false
	}
	return result
}
