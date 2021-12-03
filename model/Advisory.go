package model

type Condition func() bool

type Adivsory struct {
	AdvisoryName string
	AdvisoryType string
	Severity     int
	Condition    Condition
}
