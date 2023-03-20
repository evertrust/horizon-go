package rfc5280

import (
	"net/url"
	"os"
	"testing"

	"github.com/evertrust/horizon-go/certificates"
	"github.com/evertrust/horizon-go/http"
)

var client Client
var certsClient certificates.Client

func init() {
	var baseClient = http.Client{}
	endpoint, _ := url.Parse(os.Getenv("ENDPOINT"))
	baseClient.SetBaseUrl(*endpoint)
	baseClient.InitPasswordAuth(
		os.Getenv("APIID"),
		os.Getenv("APIKEY"),
	)
	client = Client{Http: &baseClient}
}

func TestPkcs10(t *testing.T) {
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
	csr, err := client.Pkcs10(csrPem)
	if err != nil {
		t.Error(err.Error())
	}
	t.Logf("csr: %v", csr.Dn)
}

func TestTrustChain(t *testing.T) {
	certs, _, _, err := certsClient.Search("status is valid", 0, false)
	if err != nil {
		t.Error(err.Error())
	}

	detailedCert, err := certsClient.Get(certs[0].Id)
	if err != nil {
		t.Error(err.Error())
	}

	chain, err := client.Trustchain([]byte(detailedCert.Certificate), RootToLeaf)
	if err != nil {
		t.Error(err.Error())
	}

	t.Logf("root: %v", chain[0].Dn)
}
