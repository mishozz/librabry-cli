package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type UserClient interface {
	Login(username, password string) (string, error)
	Logout(token string) (string, error)
	TakeBook(token, email, isbn string) (string, error)
	ReturnBook(token, email, isbn string) error
}

type userClient struct{}

type UserDetails struct {
	Email string `json:"email"`
}

var UnauthorizedErr = errors.New("Unauthorized")

var User UserClient = &userClient{}

func (u userClient) Login(email, password string) (string, error) {
	user := &UserDetails{
		Email: email,
	}
	jsonData, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	req, _ := http.NewRequest("POST", HOST+PORT+LIBRARY_API_V1+"login", bytes.NewBuffer(jsonData))

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

func (u userClient) TakeBook(token, email, isbn string) (string, error) {
	req, _ := http.NewRequest("POST", HOST+PORT+LIBRARY_API_V1+"users/"+email+"/"+isbn, nil)
	setAuthHeader(token, req)

	respString, err := sendRequest(req)
	if err != nil {
		return "", err
	}
	return respString, nil
}

func (u userClient) ReturnBook(token, email, isbn string) error {
	req, _ := http.NewRequest("DELETE", HOST+PORT+LIBRARY_API_V1+"users/"+email+"/"+isbn, nil)
	setAuthHeader(token, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return UnauthorizedErr
	} else if resp.StatusCode != http.StatusNoContent {
		return errors.New("Unable to return book")
	}

	return nil
}
