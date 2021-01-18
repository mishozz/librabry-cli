package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	HOST           = "http://localhost:"
	PORT           = "8080"
	LIBRARY_API_V1 = "/library/api/v1/"
)

type BookClient interface {
	GetAllBooks(token string) (string, error)
	GetBook(token, isbn string) (string, error)
	SaveBook(token, isbn, title, author string, availableUnits uint) (string, error)
	Delete(token, isbn string) error
}

type bookClient struct{}

type BookDetails struct {
	Isbn           string `json:"Isbn"`
	Title          string `json:"Title"`
	Author         string `json:"Author"`
	AvailableUnits uint   `json:"AvailableUnits" `
}

var Books BookClient = &bookClient{}

func (b bookClient) GetAllBooks(token string) (string, error) {
	req, _ := http.NewRequest("GET", HOST+PORT+LIBRARY_API_V1+"books", nil)
	setAuthHeader(token, req)

	respString, err := sendRequest(req)
	if err != nil {
		return "", err
	}
	return respString, nil
}

func (b bookClient) GetBook(token, isbn string) (string, error) {
	req, _ := http.NewRequest("GET", HOST+PORT+LIBRARY_API_V1+"books/"+isbn, nil)
	setAuthHeader(token, req)

	respString, err := sendRequest(req)
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

	req, _ := http.NewRequest("POST", HOST+PORT+LIBRARY_API_V1+"books", bytes.NewBuffer(jsonData))
	setAuthHeader(token, req)

	respString, err := sendRequest(req)
	if err != nil {
		return "", err
	}
	return respString, nil
}

func (b bookClient) Delete(token, isbn string) error {
	req, _ := http.NewRequest("DELETE", HOST+PORT+LIBRARY_API_V1+"books/"+isbn, nil)
	setAuthHeader(token, req)

	client := &http.Client{}
	resp, err := client.Do(req)
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

func sendRequest(req *http.Request) (string, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
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
