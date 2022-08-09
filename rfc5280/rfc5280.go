// Package rfc5280 provides various utilities to parse and decode data formatted according to the RFC5280, such as PKCS#10s or PKCS#12s.
package rfc5280

type SubjectAlternateName struct {
	SanType string `json:"sanType"`
	Value   string `json:"value"`
}

type CFDistinguishedName struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type CFCertificationRequest struct {
	Dn         string                 `json:"dn"`
	Sans       []SubjectAlternateName `json:"sans"`
	DnElements []CFDistinguishedName  `json:"dnElements"`
	KeyType    string                 `json:"keyType"`
	Pem        string                 `json:"pem"`
}

type CfCertificate struct {
	Dn               string                 `json:"dn"`
	Sans             []SubjectAlternateName `json:"sans"`
	DnElements       []CFDistinguishedName  `json:"dnElements"`
	KeyType          string                 `json:"keyType"`
	SigningAlgorithm string                 `json:"signingAlgorithm"`
	Pem              string                 `json:"pem"`
	Serial           string                 `json:"serial"`
	IssuerDn         string                 `json:"issuerDn"`
	NotBefore        int                    `json:"notBefore"`
	NotAfter         int                    `json:"notAfter"`
	SelfSigned       bool                   `json:"selfSigned"`
}

type TrustchainOrder int

const (
	RootToLeaf TrustchainOrder = iota
	LeafToRoot
	IssuingRootToLeaf
	IssuingLeafToRoot
)

func (order TrustchainOrder) String() string {
	switch order {
	case RootToLeaf:
		return "rtl"
	case LeafToRoot:
		return "ltr"
	case IssuingRootToLeaf:
		return "irtl"
	case IssuingLeafToRoot:
		return "iltr"
	}
	return ""
}
