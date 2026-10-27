package models

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

var ErrNoPEMCertificate = errors.New("no PEM certificate block found")

// X509 parses the PEM-encoded certificate returned by Horizon.
func (o *CFCertificateResponse) X509() (*x509.Certificate, error) {
	block, _ := pem.Decode([]byte(o.Pem))
	if block == nil || block.Type != "CERTIFICATE" {
		return nil, ErrNoPEMCertificate
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("could not parse certificate %q: %w", o.Dn, err)
	}

	return cert, nil
}
