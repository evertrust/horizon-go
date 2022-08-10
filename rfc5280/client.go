// Package rfc5280 provides utilities to interact with the Horizon api.rfc5280 APIs.
package rfc5280

import (
	"github.com/evertrust/horizon-go/http"
	baseHttp "net/http"
	"net/url"
)

type Client struct {
	Http *http.Client
}

// Pkcs10 uses the Horizon instance to parse CSRs, avoiding doing local crytographic operations.
// This should be preferred to parsing PKCS#10 locally as this allow to have a reproductible environment.
func (c *Client) Pkcs10(pkcs10 []byte) (*CFCertificationRequest, error) {
	baseUrl := c.Http.BaseUrl()
	encodedCsr := url.PathEscape(string(pkcs10))
	reqUrl := baseUrl.String() + "/api/v1/rfc5280/pkcs10/" + encodedCsr
	req, err := baseHttp.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, err
	}
	baseResponse, err := c.Http.Do(req)
	if err != nil {
		return nil, err
	}

	response, err := c.Http.Unmarshal(baseResponse)
	if err != nil {
		return nil, err
	}

	var csr CFCertificationRequest
	err = response.Json().Decode(&csr)
	if err != nil {
		return nil, err
	}
	return &csr, nil
}

// Trustchain takes an X509 certificate and returns a collection of CfCertificate objects,
// in the order given by TrustchainOrder.
func (c *Client) Trustchain(cert []byte, order TrustchainOrder) ([]CfCertificate, error) {
	baseUrl := c.Http.BaseUrl()
	encodedCert := url.PathEscape(string(cert))
	reqUrl := baseUrl.String() + "/api/v1/rfc5280/tc/" + encodedCert + "?order=" + order.String()
	req, err := baseHttp.NewRequest("GET", reqUrl, nil)

	if err != nil {
		return nil, err
	}
	baseResponse, err := c.Http.Do(req)
	if err != nil {
		return nil, err
	}

	response, err := c.Http.Unmarshal(baseResponse)
	if err != nil {
		return nil, err
	}

	if response.HasContentType("text/plain") {
		return nil, &http.NotImplementedError{
			Feature:       http.TrustchainDecoding,
			ImplementedIn: "2.2.2",
		}
	}

	var trustchain []CfCertificate
	err = response.Json().Decode(&trustchain)
	if err != nil {
		return nil, err
	}
	return trustchain, nil
}
