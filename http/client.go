// Package http provides low-level methods to interact with the Horizon instance through its REST API.
package http

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	baseClient http.Client
	Transport  http.Transport
	baseUrl    url.URL
	cert       string
	key        string
	apiId      string
	apiKey     string
}

// Init initializes the instance parameters such as its location, and authentication data.
func (c *Client) Init(baseUrl url.URL, apiId string, apiKey string, cert string, key string) {
	c.baseUrl = baseUrl
	c.apiId = apiId
	c.apiKey = apiKey
	c.cert = cert
	c.key = key
	c.Transport.TLSClientConfig = &tls.Config{
		Certificates: []tls.Certificate{},
	}
	c.baseClient.Transport = &c.Transport
}

// SetCaBundle sets the client certificate than can be used for authentication.
func (c *Client) SetCaBundle(caBundle string) {
	caCertPool, _ := x509.SystemCertPool()
	if caCertPool == nil {
		caCertPool = x509.NewCertPool()
	}
	caCertPool.AppendCertsFromPEM([]byte(caBundle))
	c.Transport.TLSClientConfig.RootCAs = caCertPool
}

func (c *Client) SkipTLSVerify() {
	c.Transport.TLSClientConfig.InsecureSkipVerify = true
}

func (c *Client) SetProxy(proxyUrl url.URL) {
	c.Transport.Proxy = http.ProxyURL(&proxyUrl)
}

func (c *Client) Unmarshal(r *http.Response) (*HorizonResponse, error) {
	var responseBodyCopy bytes.Buffer
	teeReader := io.TeeReader(r.Body, &responseBodyCopy)
	d := json.NewDecoder(teeReader)

	if r.StatusCode > 300 {
		if r.Header.Get("Content-Type") == "application/json" {
			// Deserialize the response to an error
			var horizonError HorizonErrorResponse
			var horizonMultiError HorizonMultipleErrorsResponse
			if err := d.Decode(&horizonError); err != nil {
				if err := json.Unmarshal(responseBodyCopy.Bytes(), &horizonMultiError); err != nil {
					log.Fatalf("(HTTP %d) error deserializing error JSON: %s", r.StatusCode, string(responseBodyCopy.Bytes()))
				}
				return &HorizonResponse{
					BaseResponse: r,
				}, &horizonMultiError
			}
			return &HorizonResponse{
				BaseResponse: r,
			}, &horizonError
		} else {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				body = []byte("")
			}

			return nil, &HorizonErrorResponse{
				Code:    "Unknown",
				Message: "Non-JSON error from Horizon",
				Detail:  string(body),
			}
		}
	}

	response := HorizonResponse{
		BaseResponse: r,
	}

	return &response, nil
}

func (c *Client) Get(path string) (response *HorizonResponse, err error) {
	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	baseResponse, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	return c.Unmarshal(baseResponse)
}

func (c *Client) Post(path string, body []byte) (response *HorizonResponse, err error) {
	req, err := c.newRequest("POST", path, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	baseResponse, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	return c.Unmarshal(baseResponse)
}

func (c *Client) PostWithJwt(path, jwt string, body []byte) (response *HorizonResponse, err error) {
	req, err := c.newRequest("POST", path, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-JWT-CERT-POP", jwt)

	baseResponse, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	return c.Unmarshal(baseResponse)
}

func (c *Client) GetWithJwt(path, jwt string) (response *HorizonResponse, err error) {
	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-JWT-CERT-POP", jwt)

	baseResponse, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	return c.Unmarshal(baseResponse)
}

func (c *Client) Delete(path string) (response *HorizonResponse, err error) {
	req, err := c.newRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}
	baseResponse, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	return c.Unmarshal(baseResponse)
}

func (c *Client) Put(path string, body []byte) (response *HorizonResponse, err error) {
	req, err := c.newRequest("PUT", path, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	baseResponse, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	return c.Unmarshal(baseResponse)
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("x-api-id", c.apiId)
	req.Header.Add("x-api-key", c.apiKey)
	// Always prefer JSON responses
	req.Header.Add("Accept", "application/json")
	tempCert, err := tls.LoadX509KeyPair(c.cert, c.key)
	if err != nil {
		fmt.Printf("ERROR: %s", err)
	}
	c.Transport.TLSClientConfig.Certificates = append(c.Transport.TLSClientConfig.Certificates, tempCert)
	return c.baseClient.Do(req)
}

func (c *Client) BaseUrl() url.URL {
	return c.baseUrl
}

func (c *Client) newRequest(method string, path string, body io.Reader) (*http.Request, error) {
	reqUrl := c.baseUrl.ResolveReference(&url.URL{Path: path})
	return http.NewRequest(method, reqUrl.String(), body)
}
