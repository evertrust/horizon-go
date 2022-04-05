package requests

import (
	"github.com/evertrust/horizon-go/http"
	"net/url"
	"testing"
)

var client Client

var profile = "Profile"

func init() {
	var baseClient = http.Client{}
	baseClient.Init(
		url.URL{
			Scheme: "http",
			Host:   "localhost:9000",
		},
		"administrator",
		"github.com/evertrust/horizon-go",
	)
	client = Client{Http: &baseClient}
}

// Enroll a certificate
func TestCentralizedEnroll(t *testing.T) {
	_, err := client.CentralizedEnroll(
		"Profile",
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
	)
	if err != nil {
		t.Error(err.Error())
	}
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
		profile,
		csrPem,
		[]LabelElement{},
	)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetRequest(t *testing.T) {
	request, err := client.Get("61e6de8e3200003e007af729")

	print(request.CertificatePEM)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestRevokeRequest(t *testing.T) {
	var certificate = `-----BEGIN CERTIFICATE-----
MIIEdDCCAlygAwIBAgIIXHWpuNIbkbgwDQYJKoZIhvcNAQELBQAwQzELMAkGA1UE
BhMCRlIxEjAQBgNVBAoTCUV2ZXJUcnVzdDEgMB4GA1UEAxMXRXZlclRydXN0IFFB
IElzc3VpbmcgQ0EwHhcNMjIwMTE0MTQwNjIxWhcNMjMwMTE0MTQwNjIxWjAbMRkw
FwYDVQQDDBBGb3JGdXR1cmVSZWNvdmVyMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A
MIIBCgKCAQEA6yY52rxJP7Ew7npdwY5tIcw40+K5mxD5rV90ED+BkTlM3apfUOdy
Qy3RP49l6JRW72tnSvXId7BlgBLFJWocwcQp+78GAX/lB5+/man+v8msHgQzXQ8W
41DaQ2FU7jjwRWlD5jxBcwDwgrqD2enKN9+D6PUGjjj6fh/x54Q6tCvBQSuc+lH7
xrxrNL16Za8tPSCuRIxfPgcFuJudhDb/91ab7cBvdywFaZ/63mOHIvbrujCHyu0G
21/p5bMUsZt9iIt+mzYV2DSrPWXQZrR+r6FxVfO6GbU6YVGi+NO2WpwDwnCL9SPR
PpF5R8uImok/tKCUcrmBedy+fkxBJWTBEwIDAQABo4GTMIGQMAwGA1UdEwEB/wQC
MAAwHQYDVR0OBBYEFEgAkTiUYXX8JznUiZ5Jk+7uYJU5MB8GA1UdIwQYMBaAFBQQ
3LAzCfegGprKagjWFldyJcCpMA4GA1UdDwEB/wQEAwIFoDAdBgNVHSUEFjAUBggr
BgEFBQcDAgYIKwYBBQUHAwEwEQYDVR0gBAowCDAGBgRVHSAAMA0GCSqGSIb3DQEB
CwUAA4ICAQBNnleMCWCeUl0kKAg9G3YvvufXtj4M94vfv0rr7AUFxtwgAtDZ9wVB
kIUR6HwI97yqdW5J9Fje5VNBW0dlbcZ7lyG/kOiC4vZ/ay8QshjWuJaxmofglDR0
A/sbYQ+h+7d7zhWlpwSKYyV0HWjmPqtraah/U/0LZsrV30SZvWYaicmuHeuWBj3p
hN1usgclDCpn2map0pU0F1S9+FVHAmA6CCi17VcyfMkUxXCn66b8R2tpA2JKqzJ2
7cEQrOV9UNWCPvlyIB7yvbZrZH5iaCce2xTku8xV91O4Stsd6DPAT4gshQBdaNGC
RJwNWyAE9v/Dyn2/378B9xPzgTheG4J38vTh+t7X8io36LQPvuWdEvbWKWSWARHm
KGNR1rdcROPMsp3noMbh576pUrrgfwu6DlTRvrQCxOFveGQ460ZmIyM9Lkz3Fnyy
TMYenCdzXPBBu+e4CpRdN5BvhIt7jU83j1M+28hQ9hjJzdmgWUVnR/PM3hNdRWLo
vNhK6FHA4B2qn7qhCDWeWAo1q/F935g8wyIYRf6XngznTHLLncmLko54Zs+gI8nc
IwetYvghKt4EAZTy8km9iU9QCChRQft4CCgFNUId7peM0INTSa9TuSTHB7wJR3QT
u7u5by1YMW5wSQIyN9XAUvWm4IssesaggTboXkNcBF+EyHRZA6FgZw==
-----END CERTIFICATE-----`

	_, err := client.Revoke(certificate, "UNSPECIFIED")
	//print(request.Status)
	if err != nil {
		t.Error(err.Error())
	}
}
