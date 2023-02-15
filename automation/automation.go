package automation

type Policy struct {
	KeyType    string              `json:"keyType"`
	Module     string              `json:"module"`
	Profile    string              `json:"profile"`
	TrustChain map[string][]string `json:"trustChain"`
	// EST only parameters
	AuthorizationMode string `json:"authorizationMode,omitempty"`
	// ACME only parameters
	TlsAlpn01Port int `json:"tlsAlpn01Port,omitempty"`
	Http01Port    int `json:"http01Port,omitempty"`
}
