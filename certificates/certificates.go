// Package certificates provides utilities to interact with the Horizon api.certificate APIs.
package certificates

type RevocationReason string

const (
	RevocationReasonUnspecified          RevocationReason = "UNSPECIFIED"
	RevocationReasonKeyCompromise        RevocationReason = "KEYCOMPROMISE"
	RevocationReasonCACompromise         RevocationReason = "CACOMPROMISE"
	RevocationReasonAffiliationChange    RevocationReason = "AFFILIATIONCHANGE"
	RevocationReasonSuperseded           RevocationReason = "SUPERSEDED"
	RevocationReasonCessationOfOperation RevocationReason = "CESSATIONOFOPERATION"
)

type Certificate struct {
	Module              string           `json:"module"`
	Profile             string           `json:"profile,omitempty"`
	Owner               string           `json:"owner,omitempty"`
	Certificate         string           `json:"certificate"`
	Thumbprint          string           `json:"thumbprint"`
	SelfSigned          bool             `json:"selfSigned"`
	PublicKeyThumbprint string           `json:"publicKeyThumbprint"`
	Dn                  string           `json:"dn"`
	Serial              string           `json:"serial"`
	Issuer              string           `json:"issuer"`
	NotBefore           int              `json:"notBefore"`
	NotAfter            int              `json:"notAfter"`
	RevocationDate      int              `json:"revocationDate,omitempty"`
	RevocationReason    RevocationReason `json:"revocationReason,omitempty"`
	KeyType             string           `json:"keyType"`
	SigningAlgorithm    string           `json:"signingAlgorithm"`
}
