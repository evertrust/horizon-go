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
