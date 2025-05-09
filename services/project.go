package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

// type ProjectList struct {
// 	Projects []Project
// }

// type ProjectSearchOpts struct {
// 	AnalyzedBefore    string `json:"analyzedBefore"`
// 	OnProvisionedOnly bool   `json:"onProvisionedOnly"`
// 	Page              int    `json:"p"`
// 	Projects          string `json:"projects"`
// 	PageSize          int    `json:"ps"`
// 	Query             string `json:"q"`
// }

func NewProjectClient() *ProjectClient {
	return &ProjectClient{
		Client: client.NewClient(
			os.Getenv("SONAR_HOST"),
			os.Getenv("SONAR_TOKEN"),
			nil,
		),
	}
}

func (p *ProjectClient) GetProjects(opts map[string]string) []Project {
	// TODO: this currently assumes all opts are valid
	const pageSize int = 200
	page := 1
	projects := []Project{}

	// add query params
	endpoint := p.Client.BaseURL + "/projects/search?"
	for key, val := range opts {
		endpoint = fmt.Sprintf("%s%s=%s&", endpoint, key, val)
	}

	// token
	bearer := "Bearer " + p.Client.Token

	for { // while there are more pages
		uri := fmt.Sprintf("%sp=%d", endpoint, page)
		req, err := http.NewRequest("GET", uri, nil)
		if err != nil {
			log.Panicln("Error while creating request:", err)
		}
		req.Header.Add("Authorization", bearer)

		resp, err := p.Client.Client.Do(req)
		if err != nil {
			log.Panicln("Failed to execute request:", err)
		}
		defer resp.Body.Close()

		// read response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Panicln("Failed to read body:", err)
		}

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
