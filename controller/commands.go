package controller

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
)

func GetAdvisory() components.Command {
	return components.Command{
		Name:        "advise",
		Description: "Provides advise related to security",
		Aliases:     []string{"adv"},
		Arguments:   getAdvisoryArguments(),
		Action: func(c *components.Context) error {
			return getAdvisory(c)
		},
	}
}

func getAdvisoryArguments() []components.Argument {
	return []components.Argument{
		{
			Name:        "security",
			Description: "Search for security advises",
		},
		{
			Name:        "performance",
			Description: "Search for performance advises",
		},
	}
}

func getAdvisory(c *components.Context) error {
	if len(c.Arguments) != 1 {
		return errors.New("Wrong number of arguments. Expected: 1, " + "Received: " + strconv.Itoa(len(c.Arguments)))
	}
	var advisoryType = c.Arguments[0]
	if advisoryType == "security" {
		for index, advise := range GetSecurityAdvises() {
			fmt.Println("Running condition ", index, advise.AdvisoryInfo().AdvisoryName, "Result: ", advise.Condition())
		}
	} else if advisoryType == "performance" {
		fmt.Println("Running performance check")
		for index, advise := range GetPerformanceAdvises() {
			fmt.Println("Running condition ", index, advise.AdvisoryInfo().AdvisoryName, "Result: ", advise.Condition())
		}
		return nil
	} else {
		return errors.New("Sub command not supported")
	}
	return nil
}