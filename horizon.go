// Package horizon provides high-level methods to interact with an Horizon instance.
package horizon

import (
	"net/url"

	"github.com/evertrust/horizon-go/certificates"
	"github.com/evertrust/horizon-go/http"
	"github.com/evertrust/horizon-go/license"
	"github.com/evertrust/horizon-go/requests"
	"github.com/evertrust/horizon-go/rfc5280"
)

type Horizon struct {
	Http        *http.Client
	Requests    *requests.Client
	License     *license.Client
	Rfc5280     *rfc5280.Client
	Certificate *certificates.Client
}

// Init initializes the instance parameters such as its location, and authentication data.
func (client *Horizon) Init(baseUrl url.URL, apiId string, apiKey string) {
	client.Http = &http.Client{}
	client.Http.Init(baseUrl, apiId, apiKey)
	client.Requests = &requests.Client{Http: client.Http}
	client.License = &license.Client{Http: client.Http}
	client.Rfc5280 = &rfc5280.Client{Http: client.Http}
	client.Certificate = &certificates.Client{Http: client.Http}
}
