package client

import (
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

func (c *Client) CreateRequest(
	method string, header map[string]string,
) *http.Request {
	// create request
	req, err := http.NewRequest(method, c.URL.String(), nil)
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
	req := c.CreateRequest("GET", map[string]string{})

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

func (c *Client) ResolvePostRequest() int {
	req := c.CreateRequest("POST", map[string]string{})

	resp, err := c.Client.Do(req)
	if err != nil {
		log.Panicln("Failed to execute request:", err)
	}
	defer resp.Body.Close()

	// read response
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Panicln("Failed to read body:", err)
	// }

	return resp.StatusCode
}
