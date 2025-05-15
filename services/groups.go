package services

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/bplong33/gonarqube/client"
)

type GroupClient struct {
	*client.Client
}

type Group struct {
	name string
}

type CreateGroupRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
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

// GetGroups returns all groups from within SonarQube with any supporting data
func (g *GroupClient) GetGroups() []Group {
	return nil
}

// GetMembership returns all of the users who are members of a given group
func (g *GroupClient) GetMembership(groupName string) []string {
	return nil
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
