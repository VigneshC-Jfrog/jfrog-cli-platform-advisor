package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jfrog/jfrog-cli-core/v2/common/commands"
	"github.com/jfrog/jfrog-cli-platform-advisor/common"
)

func GetPerformanceAdvises() []Adivsory {
	serverDetails, err := commands.NewCurlCommand().GetServerDetails()
	if err != nil {
		return nil
	}
	var advisoryType = "performance"
	var advisory []Adivsory
	advisory = append(advisory, Adivsory{
		AdvisoryName: "HighVirualRepositoryAssociation",
		Condition: func() bool {
			var result bool = true
			var repos []Repo
			url := serverDetails.GetArtifactoryUrl() + "api/repositories?type=virtual"
			httpRequest := &common.HttpRequest{ReqUrl: url, ReqType: "GET", AuthUser: serverDetails.GetUser(), AuthPass: serverDetails.GetPassword()}
			response := common.MakeHTTPCall(*httpRequest)
			err := json.Unmarshal(response, &repos)
			if err != nil {
				errors.New("Error unmarshalling repos")
			}
			for i := 0; i < len(repos); i++ {
				var url = serverDetails.GetArtifactoryUrl() + "api/repositories/" + repos[i].Key
				httpRequest := &common.HttpRequest{ReqUrl: url, ReqType: "GET", AuthUser: serverDetails.GetUser(), AuthPass: serverDetails.GetPassword()}
				var response = common.MakeHTTPCall(*httpRequest)
				var repo Repo
				err := json.Unmarshal(response, &repo)

				if err != nil {
					errors.New("Error unmarshalling repos")
				}
				if len(repo.Repositories) > 4 {
					result = false
					fmt.Println(strconv.Itoa(len(repo.Repositories)) + " are mapped to virtual repository " + repo.Key + " Consider splitting it.")
				} else if len(repo.Repositories) == 0 {
					fmt.Println("No repositories mapped to virtual repository " + repo.Key + " Consider adding one or delete this virtual repository")
				}
			}
			return result
		},
		AdvisoryType: advisoryType,
		Severity:     3})
	advisory = append(advisory, Adivsory{AdvisoryName: "HighStorageSummaryPerRepo", Condition: func() bool {
		var result bool = true
		var storageSummary StorageSummary
		var My_map = make(map[string]float64)
		var url = serverDetails.GetArtifactoryUrl() + "api/storageinfo"
		httpRequest := &common.HttpRequest{ReqUrl: url, ReqType: "GET", AuthUser: serverDetails.GetUser(), AuthPass: serverDetails.GetPassword()}
		var response = common.MakeHTTPCall(*httpRequest)
		err := json.Unmarshal(response, &storageSummary)
		if err != nil {
			errors.New("Error unmarshalling repos")
		}
		for i := 0; i < len(storageSummary.Repo); i++ {
			if storageSummary.Repo[i].RepoKey == "TOTAL" {
				continue
			}
			var size_type = strings.Split(storageSummary.Repo[i].UsedSpace, " ")[1]
			var used_size, _ = strconv.ParseFloat(strings.Split(storageSummary.Repo[i].UsedSpace, " ")[0], 64)
			if (size_type == "GB" && used_size > 3) || size_type == "TB" {
				result = false
				fmt.Println("Reposiroty " + storageSummary.Repo[i].RepoKey + " has " + storageSummary.Repo[i].UsedSpace + ". Please split the content to different repositories.")
				My_map[storageSummary.Repo[i].RepoKey+"("+size_type+")"] = used_size
			}
		}
		return result
	}, AdvisoryType: advisoryType, Severity: 1})
	return advisory
}
