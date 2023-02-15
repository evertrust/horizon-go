package automation

type Policy struct {
	KeyType           string              `json:"keyType"`
	Module            string              `json:"module"`
	Profile           string              `json:"profile"`
	TrustChain        map[string][]string `json:"trustChain"`
	AuthorizationMode string              `json:"authorizationMode,omitempty"`
}
