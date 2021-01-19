package client

import (
	"io/ioutil"
	"net/http"
)

// HttpClient can be used to mock out unit testing calls
type HttpClient interface {
	SendRequest(req *http.Request) (string, error)
	Do(req *http.Request) (*http.Response, error)
}

// Implements the HttpClient interface methods
type httpClient struct {
	client *http.Client
}

var Http HttpClient = &httpClient{client: &http.Client{}}

func (h httpClient) Do(req *http.Request) (*http.Response, error) {
	return h.client.Do(req)
}

func (h httpClient) SendRequest(req *http.Request) (string, error) {
	resp, err := h.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	bodyString := string(bodyBytes)

	return bodyString, nil
}
