package model

import "github.com/jfrog/jfrog-cli-core/v2/utils/config"

type AbstractAdvisory interface {
	AdvisoryInfo() AdvisoryInfo
	Condition(serverDetails *config.ServerDetails)
}

type AdvisoryInfo struct {
	AdvisoryName string
	AdvisoryType string
	Severity     int
}
