package client

import (
	"io/ioutil"
	"net/http"
)

// HTTPClient can be used to mock out unit testing calls
type HTTPClient interface {
	SendRequest(req *http.Request) (string, error)
	Do(req *http.Request) (*http.Response, error)
}

// Implements the HttpClient interface methods
type httpClient struct {
	client *http.Client
}

// HTTP is http client which can be used for mocking
var HTTP HTTPClient = &httpClient{client: &http.Client{}}

func (h httpClient) Do(req *http.Request) (*http.Response, error) {
	return h.client.Do(req)
}

// SendRequest sends http request and returns the body of the response as a string
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
