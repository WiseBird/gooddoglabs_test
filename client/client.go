package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/WiseBird/gooddoglabs_test/dal"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type restResult struct {
	Error string
	Data  json.RawMessage
}

type Client struct {
	auth   string
	url    string
	client *http.Client
}

func NewClient(url string) *Client {
	client := &http.Client{}
	return &Client{url: url, client: client}
}

func (client *Client) setAuth(username, password string) {
	client.auth = basicAuth(username, password)
}
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func (client *Client) callService(httpMethod, method string, data url.Values) ([]byte, error) {
	req, err := http.NewRequest(httpMethod, client.url+method, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+client.auth)

	resp, err := client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := &restResult{}
	err = json.Unmarshal(responseBytes, result)
	if err != nil {
		return nil, err
	}

	if result.Error != "" {
		return nil, errors.New(result.Error)
	}

	return result.Data, nil
}

func (client *Client) Login(username, password string) error {
	client.setAuth(username, password)
	_, err := client.callService("GET", "/accounts/checkAuth", nil)
	return err
}

func (client *Client) Users() ([]dal.User, error) {
	data, err := client.callService("GET", "/users", nil)
	if err != nil {
		return nil, err
	}

	users := []dal.User{}
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (client *Client) CreateUser(username, password, firstname, lastname string) error {
	values := make(url.Values)
	values.Add("username", username)
	values.Add("password", password)
	values.Add("firstname", firstname)
	values.Add("lastname", lastname)
	_, err := client.callService("POST", "/users", values)
	return err
}
