package services

import (
	"github.com/bplong33/gonarqube/client"
)

type PermissionsClient struct {
	*client.Client
}

// func (p *PermissionsClient) RemoveGroupFromAllProjects(projects []Project) []Project {
// 	var failedChange []Project
// 	// endpoint := p.URL + "/permissions/add_group?groupName=sonar-users&permission=user&projectKey="
// 	p.URL.Path = "/permissions/add_group"
//
// 	for _, proj := range projects {
// 		uri := endpoint + proj.Key
//
// 		req := p.CreateRequest("POST", map[string]string{})
// 		// req, err := http.NewRequest("POST", uri, nil)
// 		// if err != nil {
// 		// 	log.Panicln("Error while creating request:", err)
// 		// }
// 		// req.Header.Add("Authorization", "Bearer "+p.Client.Token)
// 		resp, err := p.Client.Do(req)
// 		if err != nil {
// 			log.Panicln("Failed to execute request:", err)
// 		}
// 		defer resp.Body.Close()
//
// 		if resp.StatusCode >= 300 {
// 			failedChange = append(failedChange, proj)
// 			log.Println("Failed to modify permissions on project:", proj.Key)
// 		}
// 	}
//
// 	return failedChange
// }
