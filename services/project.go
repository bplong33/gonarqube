package services

// import (
//   "github.com/bplong33/gonarqube/client
// )

type ProjectClient struct {
	Client *ClientBuilder
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

type PrivateProjectList struct {
	Projects []Projects
}

func NewProjectClient() *ProjectClient {
  p := &ProjectClient {
    Client = NewClient(),
  }
  return &p
}
