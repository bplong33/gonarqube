package client

import (
	"io"
	"log"
	"net/http"
)

type Client struct {
	BaseURL string
	Token   string
	Client  *http.Client
}

func NewClient(baseURL string, token string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return &Client{
		BaseURL: baseURL,
		Token:   token,
		Client:  httpClient,
	}
}

func ResolveGetRequest(
	uri string,
	header map[string]string,
	client *http.Client,
) []byte {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Panicln("Error while creating request:", err)
	}
	for key, val := range header {
		req.Header.Add(key, val)
	}

	resp, err := client.Do(req)
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
