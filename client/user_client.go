package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

// UserClient is an interface with methods which will call the library REST API
//
// we can use this interface to mock out unit testing calls
type UserClient interface {
	Login(username, password string) (string, error)
	Logout(token string) (string, error)
	TakeBook(token, email, isbn string) (string, error)
	ReturnBook(token, email, isbn string) error
	GetAllUsers(token string) (string, error)
	GetUser(token, email string) (string, error)
}

type userClient struct {
	client HTTPClient
}

// UserDetails holds email and password so we can parse it to json
type UserDetails struct {
	Email string `json:"email"`
}

// UnauthorizedErr is the error when the http call is unauthorized
var UnauthorizedErr = errors.New("Unauthorized")

// User is user client which can be used for mocking
var User UserClient = &userClient{
	client: HTTP,
}

func (u userClient) Login(email, password string) (string, error) {
	user := &UserDetails{
		Email: email,
	}
	jsonData, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	req, _ := http.NewRequest("POST", HOST+PORT+libraryApiV1+"login", bytes.NewBuffer(jsonData))

	respString, err := u.client.SendRequest(req)
	if err != nil {
		return "", err
	}
	return respString, nil
}

func (u userClient) Logout(token string) (string, error) {
	req, _ := http.NewRequest("POST", HOST+PORT+libraryApiV1+"logout", nil)
	setAuthHeader(token, req)

	respString, err := u.client.SendRequest(req)
	if err != nil {
		return "", err
	}
	return respString, nil
}

func (u userClient) TakeBook(token, email, isbn string) (string, error) {
	req, _ := http.NewRequest("POST", HOST+PORT+libraryApiV1+"users/"+email+"/"+isbn, nil)
	setAuthHeader(token, req)

	respString, err := u.client.SendRequest(req)
	if err != nil {
		return "", err
	}
	return respString, nil
}

func (u userClient) ReturnBook(token, email, isbn string) error {
	req, _ := http.NewRequest("DELETE", HOST+PORT+libraryApiV1+"users/"+email+"/"+isbn, nil)
	setAuthHeader(token, req)

	resp, err := u.client.Do(req)
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

func (u userClient) GetAllUsers(token string) (string, error) {
	req, _ := http.NewRequest("GET", HOST+PORT+libraryApiV1+"users", nil)
	setAuthHeader(token, req)

	respString, err := u.client.SendRequest(req)
	if err != nil {
		return "", err
	}
	return respString, nil
}

func (u userClient) GetUser(token, email string) (string, error) {
	req, _ := http.NewRequest("GET", HOST+PORT+libraryApiV1+"users/"+email, nil)
	setAuthHeader(token, req)

	respString, err := u.client.SendRequest(req)
	if err != nil {
		return "", err
	}
	return respString, nil
}
