package model

import "encoding/xml"

type StorageSummary struct {
	Repo []Repo `json:"repositoriesSummaryList"`
}

type Repo struct {
	Key          string   `json:"key"`
	Repositories []string `json:"repositories"`
	RepoKey      string   `json:"repoKey"`
	UsedSpace    string   `json:"usedSpace"`
}

type Config struct {
	XMLName  xml.Name `xml:"config"`
	Security Security `xml:"security"`
}

type Security struct {
	AnonAccess string `xml:"anonAccessEnabled"`
}

type ConditionStatus struct {
	StatusString string
	Status       bool
}

type User struct {
	Name  string
	Uri   string
	Realm string
	Email string
	Admin bool
}
