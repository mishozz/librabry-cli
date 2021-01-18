package client

import (
	"bytes"
	"net/http"

	"github.com/docker/docker-credential-helpers/client"
)

type UserClient interface {
	Login(username, password string) (string, error)
	Logout(token string) (string, error)
}

type userClient struct{}

type UserDetails struct {
	email string
}

var nativeStore = client.NewShellProgramFunc("docker-credential-secretservice")

var User UserClient = &userClient{}

func (u userClient) Login(email, password string) (string, error) {
	jsonprep := "{\"Email\":\"" + email + "\"}"
	jsonStr := []byte(jsonprep)

	req, _ := http.NewRequest("POST", HOST+PORT+LIBRARY_API_V1+"login", bytes.NewBuffer(jsonStr))

	respString, err := sendRequest(req)
	if err != nil {
		return "", err
	}
	return respString, nil
}

func (u userClient) Logout(token string) (string, error) {
	req, _ := http.NewRequest("POST", HOST+PORT+LIBRARY_API_V1+"logout", nil)
	setAuthHeader(token, req)

	respString, err := sendRequest(req)
	if err != nil {
		return "", err
	}
	return respString, nil
}
