// Package http provides low-level methods to interact with the Horizon instance through its REST API.
package http

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/evertrust/horizon-go/log"
	"io"
	gohttp "net/http"
	"net/url"
)

type Client struct {
	client      gohttp.Client
	baseUrl     string
	headerCreds *ApiCreds
	jwt         string
}

type ApiCreds struct {
	id  string
	key string
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

func (c *Client) clearAuth() {
	c.headerCreds = nil
	c.jwt = ""
	tlsConfig := c.getTlsConfig()
	tlsConfig.Certificates = nil
}

func (c *Client) SetPasswordAuth(apiId string, apiKey string) *Client {
	c.clearAuth()
	c.headerCreds = &ApiCreds{
		id:  apiId,
		key: apiKey,
	}
	return c
}

func (c *Client) SetCertAuth(cert tls.Certificate) *Client {
	c.clearAuth()
	tlsConfig := c.getTlsConfig()
	tlsConfig.Certificates = []tls.Certificate{cert}
	return c
}

func (c *Client) SetJwtAuth(jwt string) *Client {
	c.clearAuth()
	c.jwt = jwt
	return c
}

// SetCaBundle sets the client certificate than can be used for authentication.
func (c *Client) SetCaBundle(caBundle string) *Client {
	tlsConfig := c.getTlsConfig()
	if tlsConfig.RootCAs == nil {
		pool := x509.NewCertPool()
		pool.AppendCertsFromPEM([]byte(caBundle))
		tlsConfig.RootCAs = pool
	} else {
		tlsConfig.RootCAs.AppendCertsFromPEM([]byte(caBundle))
	}
	return c
}

// SkipTLSVerify skips the TLS verification.
func (c *Client) SkipTLSVerify() *Client {
	tlsConfig := c.getTlsConfig()
	tlsConfig.InsecureSkipVerify = true
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
	if c.jwt != "" {
		log.Debug("Authentication using JWT")
		request.Header.Set("X-JWT", c.jwt)
	}
	if c.headerCreds != nil {
		log.Debug("Authentication using local account " + c.headerCreds.id)
		request.Header.Set("X-API-ID", c.headerCreds.id)
		request.Header.Set("X-API-KEY", c.headerCreds.key)
	}
	if len(c.getTlsConfig().Certificates) > 0 {
		log.Debug("Authenticating using certificate")
	}
	request.Header.Set("Content-Type", "application/json")
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
