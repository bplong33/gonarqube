package client

import (
	"net/http"
	"net/url"
)

type ClientBuilder struct {
	baseURL *url.URL
	token   string
	client  *http.Client
}

type Client struct {
	Project ProjectClient
}

func NewClient(endpoint) {
}
