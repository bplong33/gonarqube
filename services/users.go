package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/bplong33/gonarqube/client"
)

type UserClient struct {
	*client.Client
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

type UserResponse struct {
	Users []User `json:"users"`
	Page  client.Paging
}

func NewUserClient(host *url.URL, token string) *UserClient {
	return &UserClient{
		Client: &client.Client{
			URL:    host,
			Token:  token,
			Client: &http.Client{},
		},
	}
}

func (u *UserClient) GetUsers(query string, inactive bool, groupId string) ([]User, error) {
	const pageSize int = 100
	page := 1
	userList := []User{}

	u.URL.Path = "/api/v2/users-management/users"

	params := url.Values{}
	params.Add("pageSize", fmt.Sprintf("%d", pageSize))
	if query != "" {
		params.Add("q", query)
	}
	if inactive {
		params.Add("active", "false")
	}
	if groupId != "" {
		params.Add("groupId", groupId)
	}

	for {
		params.Set("pageIndex", fmt.Sprintf("%d", page))
		u.URL.RawQuery = params.Encode()

		body := u.ResolveGetRequest()

		data := &UserResponse{}
		if err := json.Unmarshal([]byte(body), data); err != nil {
			return nil, err
		}
		userList = append(userList, data.Users...)

		if page*pageSize > data.Page.Total {
			break
		}
		page++
	}

	return userList, nil
}
