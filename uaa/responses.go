package uaa

import "net/http"

type Error struct {
	Error       string `json:"error"`
	Description string `json:"error_description"`
}

type Response struct {
	Proto  string
	Status string
	Header http.Header
	Body   []byte
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	Jti         string `json:"jti"`
}

type UsersResponse struct {
	Users        []User   `json:"resources"`
	StartIndex   int      `json:"startIndex"`
	ItemsPerPage int      `json:"itemsPerPage"`
	TotalResults int      `json:"totalResults"`
	Schemas      []string `json:"schemas"`
}

type User struct {
	Id                   string        `json:"id"`
	Username             string        `json:"username"`
	Name                 Name          `json:"name"`
	PhoneNumbers         []PhoneNumber `json:"phoneNumbers"`
	Emails               []Email       `json:"emails"`
	Groups               []Group       `json:"groups"`
	Approvals            []Approval    `json:"approvals"`
	Active               bool          `json:"active"`
	LastLogonTime        int           `json:"lastLogonTime"`
	PreviousLogonTime    int           `json:"previousLogonTime"`
	Verified             bool          `json:"verified"`
	Origin               string        `json:"origin"`
	ZoneId               string        `json:"zoneId"`
	PasswordLastModified string        `json:"passwordLastModified"`
	ExternalId           string        `json:"externalId"`
	Meta                 Meta          `json:"meta"`
}

type Name struct {
	FamilyName string `json:"familyName"`
	GivenName  string `json:"givenName"`
}

type PhoneNumber struct {
	Value string `json:"value"`
}

type Email struct {
	Value   string `json:"value"`
	Primary bool   `json:"primary"`
}

type Group struct {
	Value   string `json:"value"`
	Display string `json:"display"`
	Type    string `json:"type"`
}

type Approval struct {
	UserId        string `json:"userId"`
	ClientId      string `json:"clientId"`
	Scope         string `json:"scope"`
	Status        string `json:"status"`
	LastUpdatedAt string `json:"lastUpdatedAt"`
	ExpiresAt     string `json:"expiresAt"`
}

type Meta struct {
	Version      int    `json:"version"`
	LastModified string `json:"lastModified"`
	Created      string `json:"created"`
}
