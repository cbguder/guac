package uaa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type UAA struct {
	Environment Environment
	Client      *http.Client
}

func (u *UAA) GetClientToken(clientId, clientSecret string) (Token, error) {
	values := url.Values{
		"client_id":     []string{clientId},
		"client_secret": []string{clientSecret},
		"grant_type":    []string{"client_credentials"},
		"response_type": []string{"token"},
	}

	bodyReader := strings.NewReader(values.Encode())

	req, err := u.Environment.UnauthorizedRequest("POST", "/oauth/token", bodyReader)
	if err != nil {
		return Token{}, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	tokenResponse := TokenResponse{}

	err = u.performRequestAndDecode(req, &tokenResponse)
	if err != nil {
		return Token{}, err
	}

	return processTokenResponse(tokenResponse), nil
}

func (u *UAA) GetOwnerToken(clientId, clientSecret, username, password string) (Token, error) {
	values := url.Values{
		"client_id":     []string{clientId},
		"client_secret": []string{clientSecret},
		"username":      []string{username},
		"password":      []string{password},
		"grant_type":    []string{"password"},
		"response_type": []string{"token"},
	}

	bodyReader := strings.NewReader(values.Encode())

	req, err := u.Environment.UnauthorizedRequest("POST", "/oauth/token", bodyReader)
	if err != nil {
		return Token{}, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	tokenResponse := TokenResponse{}

	err = u.performRequestAndDecode(req, &tokenResponse)
	if err != nil {
		return Token{}, err
	}

	return processTokenResponse(tokenResponse), nil
}

func (u *UAA) GetUsers() ([]User, error) {
	req, err := u.Environment.AuthorizedRequest("GET", "/Users", nil)
	if err != nil {
		return nil, err
	}

	usersResponse := UsersResponse{}

	err = u.performRequestAndDecode(req, &usersResponse)
	if err != nil {
		return nil, err
	}

	return usersResponse.Users, nil
}

func (u *UAA) GetClients() ([]Client, error) {
	req, err := u.Environment.AuthorizedRequest("GET", "/oauth/clients", nil)
	if err != nil {
		return nil, err
	}

	clientsResponse := ClientsResponse{}

	err = u.performRequestAndDecode(req, &clientsResponse)
	if err != nil {
		return nil, err
	}

	return clientsResponse.Clients, nil
}

func (u *UAA) Curl(method, path string, headers http.Header, data string) (Response, error) {
	bodyReader := strings.NewReader(data)

	req, err := u.Environment.AuthorizedRequest(method, path, bodyReader)
	if err != nil {
		return Response{}, err
	}

	if data != "" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	for key, values := range headers {
		req.Header[key] = values
	}

	res, err := u.Client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Response{}, err
	}

	return Response{
		Proto:  res.Proto,
		Status: res.Status,
		Body:   body,
		Header: res.Header,
	}, nil
}

func (u *UAA) performRequestAndDecode(req *http.Request, v interface{}) error {
	res, err := u.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	if res.StatusCode != http.StatusOK {
		e := Error{}
		decoder.Decode(&e)

		if e.Description == "" {
			return fmt.Errorf("Unexpected status: %d", res.StatusCode)
		} else {
			return fmt.Errorf("Unexpected status: %d (%s)", res.StatusCode, e.Description)
		}
	}

	err = decoder.Decode(v)
	if err != nil {
		return err
	}

	return nil
}

func processTokenResponse(tokenResponse TokenResponse) Token {
	scope := strings.Split(tokenResponse.Scope, " ")

	return Token{
		AccessToken: tokenResponse.AccessToken,
		TokenType:   tokenResponse.TokenType,
		ExpiresIn:   tokenResponse.ExpiresIn,
		Scope:       scope,
		Jti:         tokenResponse.Jti,
	}
}
