package horizon

import (
	"encoding/json"
	"fmt"
)

type Principal struct {
	Identity Identity `json:"identity"`

	// Permissions The permissions of the principal
	Permissions []Permission `json:"permissions"`

	// Roles The roles of the principal
	Roles []string `json:"roles"`

	// Teams The teams of the principal
	Teams []string `json:"teams"`
}

type InternalPrincipal struct {
	Identity Identity `json:"identity"`

	// Permissions The permissions of the principal
	Permissions []Permission `json:"permissions"`

	// Roles The roles of the principal
	Roles []string `json:"roles"`

	// Teams The teams of the principal
	Teams Teams `json:"teams"`
}
type Teams []string

func (t *Teams) UnmarshalJSON(data []byte) error {
	// Try the old format: []string
	var oldFormat []string
	if err := json.Unmarshal(data, &oldFormat); err == nil {
		*t = oldFormat
		return nil
	}

	// Try the new format: []object{name:string}
	var newFormat []struct {
		Name string `json:"name"`
	}
	if err := json.Unmarshal(data, &newFormat); err == nil {
		out := make([]string, 0, len(newFormat))
		for _, item := range newFormat {
			out = append(out, item.Name)
		}
		*t = out
		return nil
	}

	// If both fail, bail out
	return fmt.Errorf("teams: unsupported JSON format: %s", string(data))
}

// Identity The principal's identity
type Identity struct {
	// Certificate The principal's certificate (in case of `X509` identity provider)
	Certificate string `json:"certificate,omitempty"`

	// Email The principal's e-mail
	Email string `json:"email,omitempty"`

	// Identifier The principal's identifier
	Identifier string `json:"identifier"`

	// IdentityProviderName The identity provider's name this principal is registered on
	IdentityProviderName string `json:"identityProviderName,omitempty"`

	// IdentityProviderType The identity provider's type this principal is registered on
	IdentityProviderType IdentityIdentityProviderType `json:"identityProviderType,omitempty"`

	// Name The principal's name
	Name string `json:"name,omitempty"`
}

// IdentityIdentityProviderType The identity provider's type this principal is registered on
type IdentityIdentityProviderType string

// Permission defines model for Permission.
type Permission struct {
	// Filter The filter to apply to the permission in the HPQL format
	Filter string `json:"filter,omitempty"`

	// Value The permission string, in the Horizon format : `<group>:<resource>:<scope>:<action>`
	Value string `json:"value"`
}
