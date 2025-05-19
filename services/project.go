package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/bplong33/gonarqube/client"
)

type ProjectClient struct {
	*client.Client
}

type ProjectResponse struct {
	Paging     client.Paging `json:"paging"`
	Components []Project     `json:"components"`
}

type Project struct {
	Key              string `json:"key"`
	Name             string `json:"name"`
	Qualifier        string `json:"qualifier"`
	Visibility       string `json:"visibility"`
	LastAnalysisDate string `json:"lastAnalysisDate"`
	Revision         string `json:"revision"`
	Managed          bool   `json:"managed"`
}

func NewProjectClient(host *url.URL, token string) *ProjectClient {
	return &ProjectClient{
		Client: &client.Client{
			URL:    host,
			Token:  token,
			Client: &http.Client{},
		},
	}
}

func (p *ProjectClient) GetProjects(params url.Values) []Project {
	const pageSize int = 200
	params.Add("ps", fmt.Sprintf("%d", pageSize))
	page := 1
	projects := []Project{}

	// add query params
	// TODO: this currently assumes all opts are valid
	p.URL.Path = "/api/projects/search"
	p.URL.RawQuery = params.Encode()

	for { // while there are more pages
		params = p.URL.Query()
		params.Set("p", fmt.Sprintf("%d", page)) // override p parameter
		p.URL.RawQuery = params.Encode()

		body := p.ResolveGetRequest()

		data := &ProjectResponse{}
		if err := json.Unmarshal([]byte(body), data); err != nil {
			log.Panicln("Error while reading response:", err)
		}

		projects = append(projects, data.Components...)

		if page*pageSize > data.Paging.Total {
			break
		}
		page++
	}

	return projects
}

// func (p *ProjectClient) {
//
// }
