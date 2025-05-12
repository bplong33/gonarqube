package services

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/bplong33/gonarqube/client"
)

//   "github.com/bplong33/gonarqube/client

type ProjectClient struct {
	Client *client.Client
}

type ProjectResponse struct {
	Paging     Paging    `json:"paging"`
	Components []Project `json:"components"`
}

type Paging struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
	Total     int `json:"total"`
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

func NewProjectClient() *ProjectClient {
	return &ProjectClient{
		Client: client.NewClient(
			os.Getenv("SONAR_HOST"),
			os.Getenv("SONAR_TOKEN"),
			nil,
		),
	}
}

func (p *ProjectClient) GetProjects(params map[string]string) []Project {
	const pageSize int = 200
	page := 1
	projects := []Project{}

	// add query params
	endpoint := p.Client.BaseURL + "/projects/search?"
	// TODO: this currently assumes all opts are valid
	for key, val := range params {
		endpoint = fmt.Sprintf("%s%s=%s&", endpoint, key, val)
	}

	// token
	bearer := "Bearer " + p.Client.Token
	header := map[string]string{"Authorization": bearer}

	for { // while there are more pages
		uri := fmt.Sprintf("%sp=%d", endpoint, page)
		body := client.ResolveGetRequest(uri, header, p.Client.Client)

		data := &ProjectResponse{}
		if err := json.Unmarshal([]byte(body), data); err != nil {
			log.Panicln("Error while reading rsponse:", err)
		}

		projects = append(projects, data.Components...)

		if page*pageSize > data.Paging.Total {
			break
		}
		page += 1
	}
	return projects
}
