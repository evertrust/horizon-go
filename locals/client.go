package locals

import (
	"encoding/json"

	"github.com/evertrust/horizon-go/http"
)

type Client struct {
	Http *http.Client
}

func (c *Client) GetAccount(identifier string) (*LocalAccount, error) {
	var local LocalAccount
	response, err := c.Http.Get("/api/v1/security/identity/locals/" + identifier)
	if err != nil {
		return nil, err
	}

	err = response.Json().Decode(&local)
	if err != nil {
		return nil, err
	}
	return &local, nil
}

func (c *Client) GetAllAccounts() (*http.HorizonResponse, error) {
	response, err := c.Http.Get("/api/v1/security/identity/locals")
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) Create(identifier string, email string) (*LocalAccount, error) {
	var local LocalAccount
	local.Identifier = identifier
	if email != "" {
		local.Email = email
	}
	jsonData, _ := json.Marshal(local)

	response, err := c.Http.Post("/api/v1/security/identity/locals", jsonData)
	if err != nil {
		return nil, err
	}
	err = response.Json().Decode(&local)
	if err != nil {
		return nil, err
	}
	return &local, nil
}

func (c *Client) Delete(acc *LocalAccount) error {
	identifier := acc.Identifier
	_, err := c.Http.Delete("/api/v1/security/identity/locals/" + identifier)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) SetPassword(acc *LocalAccount, password string) (string, error) {
	var local LocalAccount
	local.Identifier = acc.Identifier
	local.Password = password

	jsonData, _ := json.Marshal(local)
	_, err := c.Http.Patch("/api/v1/security/identity/locals", jsonData)
	if err != nil {
		return "", err
	}

	return password, nil
}

func (c *Client) AssignRoles(acc *LocalAccount, contact string, roles []string) error {
	var reqRoles []string
	reqRoles = append(reqRoles, roles...)

	var principal PrincipalInfos
	principal.Identifier = acc.Identifier
	principal.Contact = contact
	principal.Roles = append(principal.Roles, reqRoles...)

	jsonData, _ := json.Marshal(principal)

	_, err := c.Http.Post("/api/v1/security/principalinfos", jsonData)
	if err != nil {
		return err
	}

	return nil
}
