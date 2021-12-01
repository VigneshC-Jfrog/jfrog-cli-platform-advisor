package commands

import (
	"errors"
	"fmt"

	"github.com/jfrog/jfrog-cli-platform-advisor/inputs"
)

func securityAdvisory() error {
	var config_xml = inputs.GetConfig()
	if config_xml.Security.AnonAccess == "" {
		return errors.New("Cannot fetch anonymous access flag fro config.xml")
	}
	if config_xml.Security.AnonAccess == "false" {
		fmt.Println("Anonymous access is disabled. You are safe!")
		return nil
	} else {
		fmt.Println("Anonymous access is enabled. Advise to disable it if the instance is publically accessible")
		return nil
	}
}
