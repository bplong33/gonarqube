package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/bplong33/gonarqube/client"
)

type GroupClient struct {
	*client.Client
}

type Group struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Managed     bool   `json:"managed,omitempty"`
	Default     bool   `json:"default,omitempty"`
}

type CreateGroupRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type GroupResponse struct {
	Groups []Group       `json:"groups"`
	Page   client.Paging `json:"page"`
}

func NewGroupClient(host *url.URL, token string) *GroupClient {
	return &GroupClient{
		Client: &client.Client{
			URL:    host,
			Token:  token,
			Client: &http.Client{},
		},
	}
}

// CreateGroup creates a new group and returns the status of the request
func (g *GroupClient) CreateGroup(groupName string, description string) (int, error) {
	group := CreateGroupRequest{Name: groupName, Description: description}
	header := map[string]string{"Content-Type": "application/json"}
	jsonBody, err := json.Marshal(group)
	if err != nil {
		return 0, err
	}

	g.URL.Path = "/api/v2/authorizations/groups"
	statusCode, _, err := g.ResolvePostRequest(header, jsonBody)
	if err != nil {
		return 0, err
	}

	return statusCode, nil
}

// GetGroups returns all groups from within SonarQube with any supporting data
func (g *GroupClient) GetGroups(query string, managed bool) ([]Group, error) {
	const pageSize int = 100
	page := 1
	allGroups := []Group{}

	params := url.Values{}
	params.Add("q", query)
	params.Add("pageSize", fmt.Sprintf("%d", pageSize))

	g.URL.Path = "/api/v2/authorizations/groups"

	for {
		params.Set("pageIndex", fmt.Sprintf("%d", page))
		g.URL.RawQuery = params.Encode()

		body := g.ResolveGetRequest()

		data := &GroupResponse{}
		if err := json.Unmarshal([]byte(body), data); err != nil {
			return nil, err
		}

		allGroups = append(allGroups, data.Groups...)
		if page*pageSize > data.Page.Total {
			break
		}
		page++
	}

	return allGroups, nil
}

func (g *GroupClient) GetGroupDetails() *Group {
	return &Group{}
}

// GetMembership returns all of the users who are members of a given group
func (g *GroupClient) GetMembership(groupName string) []string {
	return nil
}

// AddUserToGroup will add a given user as a member of the provided group and
// returns the status code of the request
func (g *GroupClient) AddUserToGroup(userId string, groupId string) (int, error) {
	g.URL.Path = "/api/v2/authorizations/group-memberships"

	body := map[string]string{
		"userId":  userId,
		"groupId": groupId,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return -1, err
	}

	statusCode, _, err := g.ResolvePostRequest(nil, bodyBytes)
	if err != nil {
		return -1, err
	}

	return statusCode, nil
}
