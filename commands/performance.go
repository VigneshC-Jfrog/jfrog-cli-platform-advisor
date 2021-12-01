package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jfrog/jfrog-cli-core/v2/common/commands"
	"github.com/jfrog/jfrog-cli-platform-advisor/common"
)

type Repo struct {
	Key          string   `json:"key"`
	Repositories []string `json:"repositories"`
	RepoKey      string   `json:"repoKey"`
	UsedSpace    string   `json:"usedSpace"`
}

type StorageSummary struct {
	Repo []Repo `json:"repositoriesSummaryList"`
}

func performanceAdvisory() error {
	getVirtualRepositoryAdvise()
	getStorageSummaryAdvise()
	return nil

}

func getStorageSummaryAdvise() error {
	serverDetails, err := commands.NewCurlCommand().GetServerDetails()
	if err != nil {
		return err
	} else {
		var storageSummary StorageSummary
		var url = serverDetails.GetArtifactoryUrl() + "api/storageinfo"
		httpRequest := &common.HttpRequest{ReqUrl: url, ReqType: "GET", AuthUser: serverDetails.GetUser(), AuthPass: serverDetails.GetPassword()}
		var response = common.MakeHTTPCall(*httpRequest)
		if string(response) == "" {
			return errors.New("No response error")
		}
		err := json.Unmarshal(response, &storageSummary)
		if err != nil {
			errors.New("Error unmarshalling repos")
		}
		for i := 0; i < len(storageSummary.Repo); i++ {
			var size_type = strings.Split(storageSummary.Repo[i].UsedSpace, " ")[1]
			var used_size, _ = strconv.ParseFloat(strings.Split(storageSummary.Repo[i].UsedSpace, " ")[0], 64)
			if (size_type == "GB" && used_size > 5) || size_type == "TB" {
				fmt.Println("Reposiroty " + storageSummary.Repo[i].RepoKey + " has " + storageSummary.Repo[i].UsedSpace + ". Please split the content to different repositories.")
			}
		}
	}
	return nil
}

func getVirtualRepositoryAdvise() error {
	serverDetails, err := commands.NewCurlCommand().GetServerDetails()
	if err != nil {
		return err
	} else {
		var url = serverDetails.GetArtifactoryUrl() + "api/repositories?type=virtual"
		httpRequest := &common.HttpRequest{ReqUrl: url, ReqType: "GET", AuthUser: serverDetails.GetUser(), AuthPass: serverDetails.GetPassword()}
		var response = common.MakeHTTPCall(*httpRequest)
		if string(response) == "" {
			return errors.New("No response error")
		}
		var repos []Repo
		err := json.Unmarshal(response, &repos)
		if err != nil {
			errors.New("Error unmarshalling repos")
		}
		for i := 0; i < len(repos); i++ {
			var url = serverDetails.GetArtifactoryUrl() + "api/repositories/" + repos[i].Key
			httpRequest := &common.HttpRequest{ReqUrl: url, ReqType: "GET", AuthUser: serverDetails.GetUser(), AuthPass: serverDetails.GetPassword()}
			var response = common.MakeHTTPCall(*httpRequest)
			if string(response) == "" {
				return errors.New("No response error")
			}
			var repo Repo
			err := json.Unmarshal(response, &repo)
			if err != nil {
				errors.New("Error unmarshalling repos")
			}
			if len(repo.Repositories) > 2 {
				fmt.Println(strconv.Itoa(len(repo.Repositories)) + " are mapped to virtual repository " + repo.Key + " Consider splitting it.")
			} else if len(repo.Repositories) == 0 {
				fmt.Println("No repositories mapped to virtual repository " + repo.Key + " Consider adding one or delete this virtual repository")
			}
		}
		return nil
	}
}
