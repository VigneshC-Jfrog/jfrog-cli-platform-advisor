package controller

import (
	performance "github.com/jfrog/jfrog-cli-platform-advisor/advisories/performance"
	security "github.com/jfrog/jfrog-cli-platform-advisor/advisories/security"
	"github.com/jfrog/jfrog-cli-platform-advisor/model"
)

func GetSecurityAdvises() []model.AbstractAdvisory {
	return []model.AbstractAdvisory{
		security.IsAnonAccessEnabled{},
		security.IsPasswordExpirationEnabled{},
		security.IsEmailAssociatedToAllUsers{}}
}

func GetPerformanceAdvises() []model.AbstractAdvisory {
	return []model.AbstractAdvisory{
		performance.HighStorageSummaryPerRepo{},
		performance.HighVirtualRepositoryAssociation{}}
}
