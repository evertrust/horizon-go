package locals

type LocalAccount struct {
	Id              string `json:"_id,omitempty"`
	Identifier      string `json:"identifier,omitempty"`
	Password        string `json:"password"`
	Email           string `json:"email"`
	Name            string `json:"name"`
	Hash            string `json:"hash"`
	ResetUUID       string `json:"resetUUID"`
	ResetExpiration string `json:"0000-00-00T00:00:00"`
}

type PrincipalInfos struct {
	Id         string   `json:"_id,omitempty"`
	Identifier string   `json:"identifier,omitempty"`
	Contact    string   `json:"contact"`
	Roles      []string `json:"roles"`
	Teams      []string `json:"teams"`
}
