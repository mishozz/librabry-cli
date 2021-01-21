package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	host         = "http://localhost:"
	port         = "8080"
	libraryApiV1 = "/library/api/v1/"
)

// BookClient is an interface with methods which will call the library REST API
//
// we can use this interface to mock out unit testing calls
type BookClient interface {
	GetAllBooks(token string) (string, error)
	GetBook(token, isbn string) (string, error)
	SaveBook(token, isbn, title, author string, availableUnits uint) (string, error)
	Delete(token, isbn string) error
}

type bookClient struct {
	client HTTPClient
}

// BookDetails holds the details of a book so we can parse it to json
type BookDetails struct {
	Isbn           string `json:"Isbn"`
	Title          string `json:"Title"`
	Author         string `json:"Author"`
	AvailableUnits uint   `json:"AvailableUnits" `
}

// Books is book cline which can be used for mocking
var Books BookClient = &bookClient{
	client: HTTP,
}

func (b bookClient) GetAllBooks(token string) (string, error) {
	req, _ := http.NewRequest("GET", host+port+libraryApiV1+"books", nil)
	setAuthHeader(token, req)

	respString, err := b.client.SendRequest(req)
	if err != nil {
		return "", err
	}
	return respString, nil
}

func (b bookClient) GetBook(token, isbn string) (string, error) {
	req, _ := http.NewRequest("GET", host+port+libraryApiV1+"books/"+isbn, nil)
	setAuthHeader(token, req)

	respString, err := b.client.SendRequest(req)
	if err != nil {
		return "", err
	}
	return respString, nil
}
func (b bookClient) SaveBook(token, isbn, title, author string, availableUnits uint) (string, error) {
	book := &BookDetails{
		Isbn:           isbn,
		Title:          title,
		Author:         author,
		AvailableUnits: availableUnits,
	}
	jsonData, err := json.Marshal(book)
	if err != nil {
		return "", err
	}

	req, _ := http.NewRequest("POST", host+port+libraryApiV1+"books", bytes.NewBuffer(jsonData))
	setAuthHeader(token, req)

	respString, err := b.client.SendRequest(req)
	if err != nil {
		return "", err
	}
	return respString, nil
}

func (b bookClient) Delete(token, isbn string) error {
	req, _ := http.NewRequest("DELETE", host+port+libraryApiV1+"books/"+isbn, nil)
	setAuthHeader(token, req)

	resp, err := b.client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return UnauthorizedErr
	} else if resp.StatusCode != http.StatusNoContent {
		return errors.New("Unable to delete book")
	}

	return nil
}

func setAuthHeader(token string, req *http.Request) {
	tokenString := fmt.Sprintf("Bearer %v", token)
	req.Header.Set("Authorization", tokenString)
}
