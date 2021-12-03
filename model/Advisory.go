package model

type AbstractAdvisory interface {
	AdvisoryInfo() AdvisoryInfo
	Condition() bool
}

type AdvisoryInfo struct {
	AdvisoryName string
	AdvisoryType string
	Severity     int
}
