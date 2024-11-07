package client

import (
	"crypto"
	"crypto/tls"
	"crypto/x509"
	"github.com/evertrust/horizon-go/automation"
	"github.com/evertrust/horizon-go/certificates"
	"github.com/evertrust/horizon-go/discovery"
	"github.com/evertrust/horizon-go/http"
	"github.com/evertrust/horizon-go/license"
	"github.com/evertrust/horizon-go/locals"
	mylog "github.com/evertrust/horizon-go/log"
	"github.com/evertrust/horizon-go/principals"
	"github.com/evertrust/horizon-go/requests"
	"github.com/evertrust/horizon-go/rfc5280"
	"io"
	"log"
	gohttp "net/http"
	"net/url"
	"time"
)

type Client struct {
	Requests    *requests.Client
	License     *license.Client
	Rfc5280     *rfc5280.Client
	Certificate *certificates.Client
	Discovery   *discovery.Client
	Automation  *automation.Client
	Locals      *locals.Client
	Http        *http.Client
	Principals  *principals.Client
}

// New instantiates a new Client client
func New(httpClient *http.Client) *Client {
	var client Client
	if httpClient == nil {
		httpClient = &http.Client{}
		httpClient.SetHttpClient(nil)
	}
	return client.init(httpClient)
}

func (client *Client) SetDebugWriter(writer io.Writer) *Client {
	log.SetOutput(writer)
	mylog.LogEnabled = true
	return client
}

// SetHttpClient initializes the instance parameters such as its location, and authentication data.
func (c *Client) SetHttpClient(httpClient *gohttp.Client) *Client {
	c.Http.SetHttpClient(httpClient)
	return c
}

// SetBaseUrl sets the base url for the client
// This is the endpoint without any additional path
// For example: https://horizon-test.com
func (c *Client) SetBaseUrl(baseUrl url.URL) *Client {
	c.Http.SetBaseUrl(baseUrl)
	return c
}

func (c *Client) ClearAuth() {
	c.Http.ClearAuth()
}

func (c *Client) SetPasswordAuth(apiId string, apiKey string) *Client {
	c.Http.SetPasswordAuth(apiId, apiKey)
	return c
}

// SetCertAuth sets the client certificate than can be used for authentication.
func (c *Client) SetCertAuth(cert tls.Certificate) *Client {
	c.Http.SetCertAuth(cert)
	return c
}

func (c *Client) SetJwtAuth(cert x509.Certificate, key crypto.Signer) *Client {
	c.Http.SetJwtAuth(cert, key)
	return c
}

// SetCaBundle sets the CA bundle
func (c *Client) SetCaBundle(caBundle *x509.CertPool) *Client {
	c.Http.SetCaBundle(caBundle)
	return c
}

func (c *Client) GetCaBundle() *x509.CertPool {
	return c.GetTlsConfig().RootCAs
}

// SkipTLSVerify skips the TLS verification.
func (c *Client) SkipTLSVerify() *Client {
	c.Http.SkipTLSVerify()
	return c
}

// SetTimeout sets the timeout for http requests
func (c *Client) SetTimeout(timeout time.Duration) *Client {
	c.Http.SetTimeout(timeout)
	return c
}

func (c *Client) SetProxy(proxyUrl url.URL) *Client {
	c.Http.SetProxy(proxyUrl)
	return c
}

func (c *Client) GetTransport() *gohttp.Transport {
	return c.Http.GetTransport()
}

func (c *Client) GetTlsConfig() *tls.Config {
	return c.Http.GetTlsConfig()
}

func (c *Client) BaseUrl() (url.URL, error) {
	return c.Http.BaseUrl()
}

// init initializes the instance parameters such as its location, and authentication data.
func (client *Client) init(httpClient *http.Client) *Client {
	client.Http = httpClient
	client.Requests = requests.Init(httpClient)
	client.License = &license.Client{Http: client.Http}
	client.Rfc5280 = &rfc5280.Client{Http: client.Http}
	client.Certificate = certificates.Init(httpClient)
	client.Discovery = &discovery.Client{Http: client.Http}
	client.Automation = automation.Init(httpClient)
	client.Locals = &locals.Client{Http: client.Http}
	client.Principals = principals.Init(httpClient)
	return client
}
