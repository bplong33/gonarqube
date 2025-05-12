package services

import (
	"log"
	"net/http"

	"github.com/bplong33/gonarqube/client"
)

type PermissionsClient struct {
	Client *client.Client
}

func (p *PermissionsClient) RemoveGroupFromAllProjects(projects []services.Project) []services.Project {
	var failedChange []services.Project
	endpoint := p.Client.BaseURL + "/permissions/add_group?groupName=sonar-users&permission=user&projectKey="

	for _, proj := range projects {
		uri := endpoint + proj.Key

		req, err := http.NewRequest("POST", uri, nil)
		if err != nil {
			log.Panicln("Error while creating request:", err)
		}
		req.Header.Add("Authorization", "Bearer "+p.Client.Token)
		resp, err := p.Client.Client.Do(req)
		if err != nil {
			log.Panicln("Failed to execute request:", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 300 {
			failedChange = append(failedChange, proj)
			log.Println("Failed to modify permissions on project:", proj.Key)
		}
	}

	return failedChange
}
