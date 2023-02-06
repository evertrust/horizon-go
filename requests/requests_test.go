package requests

import (
	"net/url"
	"os"
	"testing"

	"github.com/evertrust/horizon-go/http"
)

var client Client

func init() {
	var baseClient = http.Client{}
	endpoint, _ := url.Parse(os.Getenv("ENDPOINT"))
	baseClient.Init(
		*endpoint,
		os.Getenv("APIID"),
		os.Getenv("APIKEY"),
		"",
		"",
	)
	client = Client{Http: &baseClient}
}

// Enroll a certificate
func TestCentralizedEnroll(t *testing.T) {
	request, err := client.CentralizedEnroll(
		os.Getenv("PROFILE"),
		"challengepassword",
		[]IndexedDNElement{
			{
				Element: "cn.1",
				Type:    "CN",
				Value:   "example.org",
			},
		},
		[]IndexedSANElement{},
		[]LabelElement{},
		"rsa-2048",
		nil,
		nil,
	)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(request.Password.Value)
}

// Sign a CSR
func TestDecentralizedEnroll(t *testing.T) {
	var csrPem = []byte(`-----BEGIN CERTIFICATE REQUEST-----
MIICvjCCAaYCAQAweTELMAkGA1UEBhMCRlIxDjAMBgNVBAgMBVBhcmlzMQ4wDAYD
VQQHDAVQYXJpczESMBAGA1UECgwJRXZlclRydXN0MRUwEwYDVQQDDAxldmVydHJ1
c3QuZnIxHzAdBgkqhkiG9w0BCQEWEGFndUBldmVydHJ1c3QuZnIwggEiMA0GCSqG
SIb3DQEBAQUAA4IBDwAwggEKAoIBAQDnnU9tevu5RWQTC1l0/FHurwS/QBzDnIqx
1VbDwMt+sFPaI1e02nvKP6iEzDU34Ub2x/SNh/jWslmpve3yNDCjJkR9TNS9YUqi
YAdARxh452njqSO1Cb9cZYBXxSbEscEWbVCFQzjXm41vChnLqjNgNKf/X+kOTrXO
jnodSehfAW4YsZ1PYgEWGX0T1BBfccQF3wJ6HunLk8/EftzDIvQge8gi1N1KtpLm
g7SMdeVnw+G8QRdl8W3J2KnSxXGX2Ip4oeIKp0Q/ItkVJ/dPktQDtguLeWP/lAXX
OY0I8BirZFVYBspkN6oz2DVG6sSplXOKWmBL0Z5Nm+g6Im5+Aj0ZAgMBAAGgADAN
BgkqhkiG9w0BAQsFAAOCAQEAcVieL6T8Nud4S71Hx3aC8Jhel1R1dDrm842Cgj0V
rYeKubzYJVWW9Eype4Q/Ydb8tNZokFRlPjej5D04Jz7eUoT+KnBbMnGTvCaPlR9/
p15duhuN9EJ1DaEfEvPiGrc38waOejJdrGXopFtTeojyi1KVwcU8EVYdaqbtuMvz
9k8zSJ1lnLB5CyPXW/TYl1pgqFd5XyQK/AgYifYDBbqdJ97xea1nN9cnyJ4NQ62F
vTvxBgHwMuplYhU1m0/KIJbhe4RTrA74wOPGS6OOZzLghcKZfQYhF6SPTeXPmGrm
VUqN/gOTLaBgj9fvEiJJFJUga4d6K+LHFW9rMhgva4GA+Q==
-----END CERTIFICATE REQUEST-----`)

	_, err := client.DecentralizedEnroll(
		os.Getenv("PROFILE"),
		csrPem,
		[]LabelElement{},
		nil,
		nil,
	)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetRequest(t *testing.T) {
	initialRequest, err := client.CentralizedEnroll(
		os.Getenv("PROFILE"),
		"challenge_password",
		[]IndexedDNElement{
			{
				Element: "cn.1",
				Type:    "CN",
				Value:   "example.org",
			},
		},
		[]IndexedSANElement{},
		[]LabelElement{},
		"rsa-2048",
		nil,
		nil,
	)
	if err != nil {
		t.Skip(err.Error())
	}

	_, err = client.Get(initialRequest.Id)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestRevokeRequest(t *testing.T) {
	initialRequest, err := client.CentralizedEnroll(
		os.Getenv("PROFILE"),
		"challenge_password",
		[]IndexedDNElement{
			{
				Element: "cn.1",
				Type:    "CN",
				Value:   "example.org",
			},
		},
		[]IndexedSANElement{},
		[]LabelElement{},
		"rsa-2048",
		nil,
		nil,
	)

	_, err = client.Revoke(initialRequest.Certificate.Certificate, "UNSPECIFIED")
	if err != nil {
		t.Error(err.Error())
	}
}
