package types

import (
	"fmt"
	"strings"
)

// Certificates

type Module string

const (
	WebRA        Module = "webra"
	Scep         Module = "scep"
	Est          Module = "est"
	Acme         Module = "acme"
	AcmeExternal Module = "acme-external"
)

type RevocationReason string

const (
	Unspecified          RevocationReason = "UNSPECIFIED"
	KeyCompromise        RevocationReason = "KEYCOMPROMISE"
	CACompromise         RevocationReason = "CACOMPROMISE"
	AffiliationChange    RevocationReason = "AFFILIATIONCHANGE"
	Superseded           RevocationReason = "SUPERSEDED"
	CessationOfOperation RevocationReason = "CESSATIONOFOPERATION"
)

func ValidateRevocationReason(reason string) (RevocationReason, error) {
	uppered := strings.ToUpper(reason)
	switch RevocationReason(uppered) {
	case Unspecified:
		return Unspecified, nil
	case KeyCompromise:
		return KeyCompromise, nil
	case CACompromise:
		return CACompromise, nil
	case AffiliationChange:
		return AffiliationChange, nil
	case Superseded:
		return Superseded, nil
	case CessationOfOperation:
		return CessationOfOperation, nil
	default:
		return "", fmt.Errorf("invalid revocation reason '%s'", reason)
	}
}

type Label struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Metadata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ThirdPartyItem struct {
	Connector   string `json:"connector"`
	Id          string `json:"id"`
	Fingerprint string `json:"fingerprint,omitempty"`
	PushDate    int    `json:"pushDate,omitempty"`
	RemoveDate  int    `json:"removeDate,omitempty"`
}

// DiscoveredTLSPorts is the struct that contains the ports
type DiscoveredTLSPorts struct {
	Port    int    `json:"port,omitempty"`
	Version string `json:"version,omitempty"`
}

// DiscoveryData is the struct of the locally discovered certs metadata
type DiscoveryData struct {
	CertificateLocation      []string             `json:"paths,omitempty"`
	CertificateUsageLocation []string             `json:"usages,omitempty"`
	OS                       []string             `json:"operatingSystems,omitempty"`
	Hostname                 []string             `json:"hostnames,omitempty"`
	IP                       string               `json:"ip,omitempty"`
	Source                   []string             `json:"sources,omitempty"`
	TLSPorts                 []DiscoveredTLSPorts `json:"tlsPorts,omitempty"`
}

type DiscoveryInfo struct {
	Campaign          string `json:"campaign"`
	LastDiscoveryDate int    `json:"lastDiscoveryDate"`
	Identifier        string `json:"identifier,omitempty"`
}

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
	ThirdPartyData        []ThirdPartyItem `json:"thirdPartyData,omitempty"`
	DiscoveryData         []DiscoveryData  `json:"discoveryData,omitempty"`
	DiscoveryInfo         []DiscoveryInfo  `json:"discoveryInfo,omitempty"`
	DiscoveryTrusted      *bool            `json:"discoveryTrusted,omitempty"`
	Labels                []Label          `json:"labels,omitempty"`
	SubjectAlternateNames []struct {
		SanType string `json:"sanType"`
		Value   string `json:"value"`
	} `json:"subjectAlternateNames"`
	Metadata []Metadata `json:"metadata"`
	HolderId string     `json:"holderId"`
}

type CertificateResponse struct {
	Certificate Certificate `json:"certificate"`
}

type CertificateSearchQuery struct {
	Query     string       `json:"query,omitempty"`
	Fields    []string     `json:"fields,omitempty"`
	SortedBy  []SortFields `json:"sortedBy,omitempty"`
	PageIndex int          `json:"pageIndex,omitempty"`
	PageSize  int          `json:"pageSize,omitempty"`
	WithCount bool         `json:"withCount,omitempty"`
}

type CertificateSearchResult struct {
	Id          string `json:"_id,omitempty"`
	Module      string `json:"module"`
	Dn          string `json:"dn"`
	Serial      string `json:"serial"`
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
