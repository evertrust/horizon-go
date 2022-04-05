// Package certificateprofiles provides utilities to interact with the Horizon api.certificate.profile APIs.
package certificateprofiles

type CertificateProfileCryptoPolicy struct {
	Centralized   bool `json:"centralized"`
	Decentralized bool `json:"decentralized"`
	Escrow        bool `json:"escrow"`
}
