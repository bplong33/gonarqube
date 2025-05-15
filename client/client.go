package client

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	URL    *url.URL
	Token  string
	Client *http.Client
}

// Paging represents the pagination data returned from SonarQube requests.
type Paging struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
	Total     int `json:"total"`
}

func (c *Client) CreateRequest(
	method string, header map[string]string, body []byte,
) *http.Request {
	// create request
	var bodyReader *bytes.Buffer = nil
	if body != nil {
		bodyReader = bytes.NewBuffer(body)
	} else {
		bodyReader = bytes.NewBuffer([]byte{}) // create an empty bytes.Buffer
	}

	req, err := http.NewRequest(method, c.URL.String(), bodyReader)
	if err != nil {
		log.Panicln("Error while creating request:", err)
	}

	// add headers
	bearer := "Bearer " + c.Token
	req.Header.Add("Authorization", bearer)
	for key, val := range header {
		req.Header.Add(key, val)
	}

	return req
}

func (c *Client) ResolveGetRequest() []byte {
	req := c.CreateRequest("GET", nil, nil)

	resp, err := c.Client.Do(req)
	if err != nil {
		log.Panicln("Failed to execute request:", err)
	}
	defer resp.Body.Close()

	// read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("Failed to read body:", err)
	}

	return body
}

func (c *Client) ResolvePostRequest(header map[string]string, body []byte) (int, string, error) {
	req := c.CreateRequest("POST", header, body)

	resp, err := c.Client.Do(req)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()

	// read response
	// respBody, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Panicln("Failed to read body:", err)
	// }

	return resp.StatusCode, resp.Status, nil
}
