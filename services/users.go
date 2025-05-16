package services

import (
	"github.com/bplong33/gonarqube/client"
)

type UserResponse struct {
	Users []User `json:"users"`
	Page  client.Paging
}

type User struct {
	Id                          string   `json:"id"`
	Login                       string   `json:"login"`
	Name                        string   `json:"name"`
	Email                       string   `json:"email"`
	Active                      bool     `json:"active"`
	Local                       bool     `json:"local"`
	Managed                     bool     `json:"managed"`
	ExternalLogin               string   `json:"exernalLogin"`
	ExternalProvider            string   `json:"externalProvider"`
	Avatar                      string   `json:"avatar"`
	SonarQubeLastConnectionDate string   `json:"sonarQubeLastConnectionDate"`
	SonarLintLastConnectionDate string   `json:"sonarLintLastConnectionDate"`
	ScmAccounts                 []string `json:"scmAccounts"`
}
