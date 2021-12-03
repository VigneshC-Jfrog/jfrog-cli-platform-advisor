package model

type StorageSummary struct {
	Repo []Repo `json:"repositoriesSummaryList"`
}

type Repo struct {
	Key          string   `json:"key"`
	Repositories []string `json:"repositories"`
	RepoKey      string   `json:"repoKey"`
	UsedSpace    string   `json:"usedSpace"`
}
