// Package http provides low-level methods to interact with the Horizon instance through its REST API.
package http

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/evertrust/horizon-go/log"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/cryptobyte"
	asn1Crypto "golang.org/x/crypto/cryptobyte/asn1"
	"io"
	"math/big"
	gohttp "net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

type Client struct {
	client      gohttp.Client
	baseUrl     string
	headerCreds *ApiCreds
	jwt         *JwtParams
}

type ApiCreds struct {
	id  string
	key string
}

type JwtParams struct {
	cert x509.Certificate
	key  crypto.Signer
}

// SetHttpClient initializes the instance parameters such as its location, and authentication data.
func (c *Client) SetHttpClient(httpClient *gohttp.Client) *Client {
	if httpClient == nil {
		httpClient = gohttp.DefaultClient
	}
	c.client = *httpClient
	return c
}

// SetBaseUrl sets the base url for the client
// This is the endpoint without any additional path
// For example: https://horizon-test.com
func (c *Client) SetBaseUrl(baseUrl url.URL) *Client {
	c.baseUrl = baseUrl.String()
	return c
}

func (c *Client) ClearAuth() {
	c.headerCreds = nil
	c.jwt = nil
	tlsConfig := c.getTlsConfig()
	tlsConfig.Certificates = nil
}

func (c *Client) SetPasswordAuth(apiId string, apiKey string) *Client {
	c.ClearAuth()
	c.headerCreds = &ApiCreds{
		id:  apiId,
		key: apiKey,
	}
	return c
}

// SetCertAuth sets the client certificate than can be used for authentication.
func (c *Client) SetCertAuth(cert tls.Certificate) *Client {
	c.ClearAuth()
	tlsConfig := c.getTlsConfig()
	tlsConfig.Certificates = []tls.Certificate{cert}
	return c
}

func (c *Client) SetJwtAuth(cert x509.Certificate, key crypto.Signer) *Client {
	c.ClearAuth()
	c.jwt = new(JwtParams)
	c.jwt.cert = cert
	c.jwt.key = key
	return c
}

func (c *Client) JwtEnabled() bool {
	return c.jwt != nil
}

func (c *Client) computeJwt(nonce string) (string, error) {
	jwt, err := computeJwtForNonce(c.jwt.cert, c.jwt.key, nonce)
	if err != nil {
		return "", fmt.Errorf("could not compute jwt: %s", err.Error())
	} else {
		return jwt, nil
	}
}

func computeJwtForNonce(cert x509.Certificate, key crypto.Signer, nonce string) (string, error) {
	claims :=
		jwt.MapClaims{
			"sub": string(pem.EncodeToMemory(&pem.Block{
				Type:  "CERTIFICATE",
				Bytes: cert.Raw,
			})),
			"iat": time.Now().Unix(),
			"exp": time.Now().Add(5 * time.Second).Unix(),
		}
	if nonce != "" {
		claims["nonce"] = nonce
	}
	signingMethod := "RS256"
	if reflect.TypeOf(cert.PublicKey) == reflect.TypeOf(&ecdsa.PublicKey{}) {
		ecdsaPublicKey := cert.PublicKey.(*ecdsa.PublicKey)
		signingMethod = fmt.Sprintf("ES%v", ecdsaPublicKey.Curve.Params().BitSize)
	}
	t := jwt.NewWithClaims(jwt.GetSigningMethod(signingMethod), &claims)
	sstr, err := t.SigningString()
	if err != nil {
		return "", err
	}
	var sig string
	// We are not signing using t.SignedString() as we need to sign the content using crypto signer.sign and not using a key
	switch signingMethod {
	case "RS256":
		if !crypto.SHA256.Available() {
			return "", errors.New("SHA256 not available")
		}

		hasher := crypto.SHA256.New()
		hasher.Write([]byte(sstr))

		// Sign the string and return the encoded bytes
		if sigBytes, err := key.Sign(rand.Reader, hasher.Sum(nil), crypto.SHA256); err == nil {
			sig = base64.RawURLEncoding.EncodeToString(sigBytes)
		} else {
			return "", err
		}
	default:
		// ECDSA
		ecdsaSigningMethod := jwt.GetSigningMethod(signingMethod).(*jwt.SigningMethodECDSA)
		if !ecdsaSigningMethod.Hash.Available() {
			return "", fmt.Errorf("hash %s not available", ecdsaSigningMethod.Hash.String())
		}

		hasher := ecdsaSigningMethod.Hash.New()
		hasher.Write([]byte(sstr))

		// Sign the string and return the encoded bytes
		if sigBytes, err := key.Sign(rand.Reader, hasher.Sum(nil), ecdsaSigningMethod.Hash); err == nil {
			// This is extracted from Sign from package ecdsa, ecdsa_legacy.go, as we need the r and s from the signature
			r, s := new(big.Int), new(big.Int)
			var inner cryptobyte.String
			input := cryptobyte.String(sigBytes)
			if !input.ReadASN1(&inner, asn1Crypto.SEQUENCE) ||
				!input.Empty() ||
				!inner.ReadASN1Integer(r) ||
				!inner.ReadASN1Integer(s) ||
				!inner.Empty() {
				return "", errors.New("invalid ASN.1 from SignASN1")
			}

			// This is extracted from SigningMethodECDSA from jwt package in Sign method in ecdsa.go
			curveBits := cert.PublicKey.(*ecdsa.PublicKey).Curve.Params().BitSize

			keyBytes := curveBits / 8
			if curveBits%8 > 0 {
				keyBytes += 1
			}

			// We serialize the outputs (r and s) into big-endian byte arrays
			// padded with zeros on the left to make sure the sizes work out.
			// Output must be 2*keyBytes long.
			out := make([]byte, 2*keyBytes)
			r.FillBytes(out[0:keyBytes]) // r is assigned to the first half of output.
			s.FillBytes(out[keyBytes:])  // s is assigned to the second half of output.
			sig = base64.RawURLEncoding.EncodeToString(out)
		} else {
			return "", err
		}
	}
	jwtstr := strings.Join([]string{sstr, sig}, ".")
	return jwtstr, nil
}

// SetCaBundle sets the CA bundle
func (c *Client) SetCaBundle(caBundle *x509.CertPool) *Client {
	tlsConfig := c.getTlsConfig()
	tlsConfig.RootCAs = caBundle
	return c
}

func (c *Client) GetCaBundle() *x509.CertPool {
	return c.getTlsConfig().RootCAs
}

// SkipTLSVerify skips the TLS verification.
func (c *Client) SkipTLSVerify() *Client {
	tlsConfig := c.getTlsConfig()
	tlsConfig.InsecureSkipVerify = true
	return c
}

// SetTimeout sets the timeout for http requests
func (c *Client) SetTimeout(timeout time.Duration) *Client {
	c.client.Timeout = timeout
	return c
}

func (c *Client) getTransport() *gohttp.Transport {
	if c.client.Transport == nil {
		tr := &gohttp.Transport{}
		c.client.Transport = tr
	}
	transportTlS := c.client.Transport.(*gohttp.Transport)
	if transportTlS.TLSClientConfig == nil {
		transportTlS.TLSClientConfig = &tls.Config{}
		c.client.Transport = transportTlS
	}
	return transportTlS
}

func (c *Client) getTlsConfig() *tls.Config {
	return c.getTransport().TLSClientConfig
}

func (c *Client) SetProxy(proxyUrl url.URL) *Client {
	transport := c.getTransport()
	transport.Proxy = gohttp.ProxyURL(&proxyUrl)
	return c
}

func (c *Client) BaseUrl() (url.URL, error) {
	baseUrl, err := url.Parse(c.baseUrl)
	return *baseUrl, err
}

func (c *Client) Get(path string) (response *HorizonResponse, err error) {
	resp, err := c.sendRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	return c.unmarshal(resp)
}

func (c *Client) Post(path string, body []byte) (response *HorizonResponse, err error) {
	resp, err := c.sendRequest("POST", path, body)
	if err != nil {
		return nil, err
	}
	return c.unmarshal(resp)
}

func (c *Client) Delete(path string) (response *HorizonResponse, err error) {
	resp, err := c.sendRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	return c.unmarshal(resp)
}

func (c *Client) Put(path string, body []byte) (response *HorizonResponse, err error) {
	resp, err := c.sendRequest("PUT", path, body)
	if err != nil {
		return nil, err
	}

	return c.unmarshal(resp)
}

func (c *Client) Patch(path string, body []byte) (response *HorizonResponse, err error) {
	resp, err := c.sendRequest("PATCH", path, body)
	if err != nil {
		return nil, err
	}

	return c.unmarshal(resp)
}

func (c *Client) sendRequest(method, urlToRequest string, body []byte) (*gohttp.Response, error) {
	// Setup url
	urlToSend, err := url.JoinPath(c.baseUrl, urlToRequest)
	if err != nil {
		return nil, err
	}
	var reader io.Reader
	if body != nil {
		reader = bytes.NewReader(body)
	}
	// Setup request
	request, err := gohttp.NewRequest(method, urlToSend, reader)
	if err != nil {
		return nil, err
	}
	// Define headers
	request.Header.Set("Content-Type", "application/json")

	// Define auth
	if c.JwtEnabled() {
		log.Debug("Authentication using JWT")
		// Do a first request to get the nonce
		requestForNonce, err := gohttp.NewRequest(method, urlToSend, strings.NewReader("{}"))
		if err != nil {
			return nil, err
		}
		requestForNonce.Header.Set("X-JWT-CERT-POP", "{}")
		requestForNonce.Header.Set("Content-Type", "application/json")
		nonceResp, err := c.client.Do(requestForNonce)
		if err != nil {
			return nil, fmt.Errorf("could not get nonce for JWT: %s", err.Error())
		}
		replayNonceHeader := nonceResp.Header.Get("Replay-Nonce")
		if replayNonceHeader == "" {
			return nil, errors.New("could not get nonce for JWT: missing Replay-Nonce header")
		}
		jwtValue, err := c.computeJwt(replayNonceHeader)
		if err != nil {
			return nil, fmt.Errorf("could not compute JWT: %s", err.Error())
		}
		request.Header.Set("X-JWT-CERT-POP", jwtValue)
	}
	if c.headerCreds != nil {
		log.Debug("Authentication using local account " + c.headerCreds.id)
		request.Header.Set("X-API-ID", c.headerCreds.id)
		request.Header.Set("X-API-KEY", c.headerCreds.key)
	}
	if len(c.getTlsConfig().Certificates) > 0 {
		log.Debug("Authenticating using certificate")
	}
	return c.client.Do(request)
}

func (c *Client) unmarshal(r *gohttp.Response) (*HorizonResponse, error) {

	if r.StatusCode > 300 {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		if r.Header.Get("Content-Type") == "application/json" || r.Header.Get("Content-Type") == "application/problem+json" {
			// Deserialize the response to an error
			var horizonError HorizonErrorResponse
			var horizonMultiError HorizonMultipleErrorsResponse
			if err := json.Unmarshal(body, &horizonError); err != nil {
				if err := json.Unmarshal(body, &horizonMultiError); err != nil {
					return nil, fmt.Errorf("cannot deserialize error JSON: %s", string(body))
				}
				return &HorizonResponse{
					HttpResponse: r,
				}, &horizonMultiError
			}
			return &HorizonResponse{
				HttpResponse: r,
			}, &horizonError
		} else {
			return nil, &HorizonErrorResponse{
				Code:    "Unknown",
				Message: "Non-JSON error from Horizon",
				Detail:  string(body),
			}
		}
	}

	response := HorizonResponse{
		HttpResponse: r,
	}

	return &response, nil
}
