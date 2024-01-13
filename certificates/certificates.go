// Package certificates provides utilities to interact with the Horizon api.certificate APIs.
package certificates

type RevocationReason string

const (
	Unspecified          RevocationReason = "UNSPECIFIED"
	KeyCompromise        RevocationReason = "KEYCOMPROMISE"
	CACompromise         RevocationReason = "CACOMPROMISE"
	AffiliationChange    RevocationReason = "AFFILIATIONCHANGE"
	Superseded           RevocationReason = "SUPERSEDED"
	CessationOfOperation RevocationReason = "CESSATIONOFOPERATION"
)

type Certificate struct {
	Id                    string           `json:"_id,omitempty"`
	Module                string           `json:"module"`
	Profile               string           `json:"profile,omitempty"`
	Owner                 string           `json:"owner,omitempty"`
	Certificate           string           `json:"certificate"`
	Thumbprint            string           `json:"thumbprint"`
	SelfSigned            bool             `json:"selfSigned"`
	PublicKeyThumbprint   string           `json:"publicKeyThumbprint"`
	Dn                    string           `json:"dn"`
	Serial                string           `json:"serial"`
	Issuer                string           `json:"issuer"`
	NotBefore             int              `json:"notBefore"`
	NotAfter              int              `json:"notAfter"`
	RevocationDate        int              `json:"revocationDate,omitempty"`
	RevocationReason      RevocationReason `json:"revocationReason,omitempty"`
	KeyType               string           `json:"keyType"`
	SigningAlgorithm      string           `json:"signingAlgorithm"`
	Revoked               bool             `json:"revoked"`
	SubjectAlternateNames []struct {
		SanType string `json:"sanType"`
		Value   string `json:"value"`
	} `json:"subjectAlternateNames"`
	Metadata []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"metadata"`
	HolderId string `json:"holderId"`
}

type HrzSearchQuery struct {
	Query     string   `json:"query,omitempty"`
	WithCount bool     `json:"withCount,omitempty"`
	PageIndex int      `json:"pageIndex,omitempty"`
	PageSize  int      `json:"pageSize,omitempty"`
	SortedBy  []string `json:"sortedBy,omitempty"`
	Fields    []string `json:"fields,omitempty"`
}

type certificateResponse struct {
	Certificate Certificate `json:"certificate"`
}

type SearchResult struct {
	Id          string `json:"_id,omitempty"`
	Module      string `json:"module"`
	DN          string `json:"dn"`
	NotAfter    int    `json:"notAfter"`
	Permissions struct {
		Revoke         bool `json:"revoke"`
		RequestRevoke  bool `json:"requestRevoke"`
		Update         bool `json:"update"`
		RequestUpdate  bool `json:"requestUpdate"`
		Recover        bool `json:"recover"`
		RequestRecover bool `json:"requestRecover"`
		Migrate        bool `json:"migrate"`
		RequestMigrate bool `json:"requestMigrate"`
	} `json:"permissions"`
}

type SearchResponse struct {
	Results   []SearchResult `json:"results"`
	PageIndex int            `json:"pageIndex"`
	PageSize  int            `json:"pageSize"`
	Count     int            `json:"count"`
	HasMore   bool           `json:"hasMore"`
}

type HorizonRequestValue struct {
	Label string `json:"label,omitempty"`
	Value string `json:"value,omitempty"`
}

type HorizonRequestTemplate struct {
	RevocationReason string                `json:"revocationReason,omitempty"`
	Team             *HorizonRequestValue  `json:"team,omitempty"`
	Owner            *HorizonRequestValue  `json:"owner,omitempty"`
	Labels           []HorizonRequestValue `json:"labels,omitempty"`
}

// HorizonRequest is a type defining an Horizon API request
type HorizonRequest struct {
	CertificatePEM string                 `json:"certificatePem,omitempty"`
	CertificateId  string                 `json:"certificateId,omitempty"`
	Workflow       string                 `json:"workflow,omitempty"`
	Profile        string                 `json:"profile,omitempty"`
	Template       HorizonRequestTemplate `json:"template,omitempty"`
}
