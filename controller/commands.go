package controller

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	sgr "github.com/foize/go.sgr"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/jfrog/jfrog-cli-core/v2/utils/config"
	"github.com/jfrog/jfrog-cli-platform-advisor/common"
	"github.com/jfrog/jfrog-cli-platform-advisor/model"
	"github.com/pterm/pterm"
)

var advisory_map = map[string][]model.AbstractAdvisory{
	"performance": GetPerformanceAdvises(),
	"security":    GetSecurityAdvises(),
	"all":         append(GetPerformanceAdvises(), GetSecurityAdvises()...)}

func GetAdvisory() components.Command {
	return components.Command{
		Name:        "advise",
		Description: "Provides advise related to security",
		Aliases:     []string{"adv"},
		Arguments:   getAdvisoryArguments(),
		Flags:       getCommonFlags(),
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
		{
			Name:        "all",
			Description: "Search for all advises",
		},
	}
}

func getCommonFlags() []components.Flag {
	return []components.Flag{
		components.StringFlag{
			Name:         "server-id",
			Description:  "Server ID for which advisory is sought.",
			DefaultValue: "",
		},
	}
}

func getAdvisory(c *components.Context) error {
	printBanner()
	if len(c.Arguments) != 1 {
		return errors.New("Wrong number of arguments. Expected: 1, " + "Received: " + strconv.Itoa(len(c.Arguments)))
	}
	var advisoryType = c.Arguments[0]
	serverIdValue := c.GetStringFlagValue("server-id")
	serverDetails, _ := config.GetSpecificConfig(serverIdValue, true, true)
	if advisory_map[advisoryType] == nil {
		return errors.New("Sub command not supported")
	}
	fmt.Printf(sgr.MustParseln("[underline]%s"), strings.ToUpper(advisoryType)+" REPORTS\n")
	for _, advise := range advisory_map[advisoryType] {
		common.AddConsoleMsg("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++\n", "default")
		common.AddConsoleMsg("Condition : "+advise.AdvisoryInfo().AdvisoryName+" Type: "+advise.AdvisoryInfo().AdvisoryType+"\n", "info")
		advise.Condition(serverDetails)
		common.AddConsoleMsg("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++\n", "default")
	}
	return nil
}

func printBanner() {
	println("")
	println("")
	pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("JFROG ", pterm.NewStyle(pterm.FgGreen)),
		pterm.NewLettersFromStringWithStyle("PLATFORM ADVISOR", pterm.NewStyle(pterm.FgGreen))).
		Render()
	println("")
	println("")
}
