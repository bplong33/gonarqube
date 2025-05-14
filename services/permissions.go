package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/bplong33/gonarqube/client"
)

type PermissionClient struct {
	*client.Client
}

type TemplateGroupResponse struct {
	Paging client.Paging     `json:"paging"`
	Groups []PermissionGroup `json:"groups"`
}

type PermissionGroup struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions"`
}

func NewPermissionClient(host *url.URL, token string) *PermissionClient {
	return &PermissionClient{
		Client: &client.Client{
			URL:    host,
			Token:  token,
			Client: &http.Client{},
		},
	}
}

func (p *PermissionClient) GetDefaultTemplate() ([]PermissionGroup, error) {
	const pageSize int = 100
	page := 1
	params := url.Values{}
	params.Add("ps", fmt.Sprintf("%d", pageSize))
	params.Add("templateName", "Default template")
	p.URL.Path = "/api/permissions/template_groups"
	p.URL.RawQuery = params.Encode()
	templateGroups := []PermissionGroup{}

	for { // while there are more pages
		params = p.URL.Query()
		params.Set("p", fmt.Sprintf("%d", page))
		p.URL.RawQuery = params.Encode()

		body := p.ResolveGetRequest()

		data := &TemplateGroupResponse{}
		if err := json.Unmarshal([]byte(body), data); err != nil {
			return nil, fmt.Errorf("Error while reading response: %w", err)
		}
		templateGroups = append(templateGroups, data.Groups...)

		if data.Paging.Total < page*pageSize {
			break
		}
		page += 1
	}

	return templateGroups, nil
}

// BulkApplyTemplate will apply a given template to all projects matching the filters.
// These filters may include visibility, projects, or query (for project)
func (p *PermissionClient) BulkApplyTemplate(params url.Values) (int, string) {
	if !params.Has("templateName") {
		params.Add("templateName", "Default template")
	}
	p.URL.Path = "/api/permissions/bulk_apply_template"
	p.URL.RawQuery = params.Encode()

	statusCode, status := p.ResolvePostRequest()

	return statusCode, status
}

// projects []Project, group string, permission string,

// BulkModifyPermission adds or removes the given permission a group on each
// project that matches the searching parameters. Return project keys that
// failed to be updated
func (p *PermissionClient) BulkModifyPermission(
	action string, group string, permission string, visibility string,
	projQuery string, projFilter string,
) ([]Project, error) {
	if !slices.Contains([]string{"add", "remove"}, action) {
		return nil, errors.New("`action` should be either `add` or `remove`")
	}
	var failedChange []Project

	// get projects matching query
	urlcp := *p.URL
	c := NewProjectClient(&urlcp, p.Token)
	getProjParams := url.Values{}
	if visibility != "" {
		getProjParams.Add("visibility", visibility)
	}
	if projQuery != "" {
		getProjParams.Add("q", projQuery)
	}
	if projFilter != "" {
		getProjParams.Add("projects", projFilter)
	}

	projectList := c.GetProjects(getProjParams)

	// Build URI
	p.URL.Path = fmt.Sprintf("/api/permissions/%s_group", action)
	params := url.Values{}
	params.Add("groupName", group)
	params.Add("permission", permission)

	numModified := 0

	for _, proj := range projectList {
		params.Set("projectKey", proj.Key)
		p.URL.RawQuery = params.Encode()

		statusCode, _ := p.ResolvePostRequest()

		if statusCode >= 300 {
			failedChange = append(failedChange, proj)
		} else {
			numModified++
		}
	}

	return failedChange, nil
}
