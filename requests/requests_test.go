package requests

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"github.com/evertrust/horizon-go/certificates"
	"github.com/evertrust/horizon-go/log"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/evertrust/horizon-go/http"
)

var client Client

func init() {
	endpoint, _ := url.Parse("https://horizon-qa.int.evertrust.fr")
	var baseClient = http.Client{}
	log.LogEnabled = true
	baseClient.SetHttpClient(nil).
		SkipTLSVerify().
		SetBaseUrl(*endpoint).
		SetPasswordAuth(
			"sma",
			"test",
		)
	client = Client{http: &baseClient}
}

// Enroll a certificate
func TestCentralizedEnroll(t *testing.T) {
	request, err := client.NewEnrollRequest(WebRAEnrollRequestParams{
		Profile: "webra-centralized",
		Template: &WebRAEnrollTemplate{
			KeyType: "rsa-2048",
			Subject: []IndexedDNElement{
				{
					Element: "cn.1",
					Type:    "CN",
					Value:   "example.org",
				},
			},
		},
		Password: "challengepassword",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(request.Password.Value)
}

func TestSearch(t *testing.T) {
	results, err := client.Search(SearchQuery{
		Fields:    []string{"profile"},
		WithCount: true,
		Scope:     "self",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(results)
}
func GenerateKeySafely(keyType string) (crypto.PrivateKey, error) {
	if !strings.Contains(keyType, "-") {
		return nil, errors.New("invalid key type: " + keyType)
	}
	keyProps := strings.Split(keyType, "-")
	if keyProps[0] != "rsa" && keyProps[0] != "ec" {
		return nil, errors.New("only RSA and ECDSA keys are supported. Invalid key type: " + keyType)
	}
	var privatekey crypto.PrivateKey
	var err error
	if keyProps[0] == "rsa" {
		keySize, err := strconv.Atoi(keyProps[1])
		if err != nil {
			return nil, errors.New("invalid key size: " + err.Error())
		}
		privatekey, err = rsa.GenerateKey(rand.Reader, keySize)
		if err != nil {
			return nil, errors.New("could not generate private key: " + err.Error())
		}
	} else if keyProps[0] == "ec" {
		var curve elliptic.Curve
		if keyProps[1] == "secp384r1" {
			curve = elliptic.P384()
		}
		if keyProps[1] == "secp256r1" {
			curve = elliptic.P256()
		}
		if curve == nil {
			return nil, errors.New("invalid curve: " + keyProps[1])
		}
		privatekey, err = ecdsa.GenerateKey(curve, rand.Reader)
		if err != nil {
			return nil, errors.New("could not generate private key: " + err.Error())
		}
	} else {
		return nil, errors.New("invalid curve type: " + keyProps[0])
	}
	return privatekey, nil
}

func CreateCsrPem(key crypto.PrivateKey) (string, error) {
	template := x509.CertificateRequest{
		Subject: pkix.Name{
			CommonName:         "CN",
			OrganizationalUnit: []string{"DEV"},
		},
		DNSNames: []string{"test.evertrust"},
	}
	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, &template, key)
	if err != nil {
		return "", errors.New("could not generate CSR: " + err.Error())
	}
	csrBlock := &pem.Block{
		Type:  "CERTIFICATE REQUEST",
		Bytes: csrBytes,
	}
	return string(pem.EncodeToMemory(csrBlock)), nil
}

func TestPopUpdate(t *testing.T) {
	// Generate CSR locally and enroll it
	key, err := GenerateKeySafely("rsa-2048")
	if err != nil {
		t.Fatal(err.Error())
	}
	csr, err := CreateCsrPem(key)
	if err != nil {
		t.Fatal(err.Error())
	}
	template, err := client.GetEnrollTemplate(WebRAEnrollTemplateParams{
		Csr:     csr,
		Profile: "webra-decentralized",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	template.KeyType = "rsa-3072"
	template.Team.Value = &HorizonString{"frontend"}
	request, err := client.NewEnrollRequest(WebRAEnrollRequestParams{
		Profile:  "webra-decentralized",
		Template: template,
		Password: "challengepassword",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	certPem := request.Certificate.Certificate
	block, _ := pem.Decode([]byte(certPem))
	if block == nil {
		t.Fatal("failed to parse certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	// Set it as JWT pop
	client.http.SetJwtAuth(*cert, key.(crypto.Signer))
	templateUpdate, err := client.GetUpdateTemplate(WebRAUpdateTemplateParams{})
	if err != nil {
		t.Fatal(err.Error())
	}
	templateUpdate.Team.Value = Delete
	response, err := client.NewUpdateRequest(WebRAUpdateRequestParams{
		Template: templateUpdate,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(templateUpdate)
	t.Log(response)
}

func TestTemplateAndCentralizedEnroll(t *testing.T) {
	template, err := client.GetEnrollTemplate(WebRAEnrollTemplateParams{
		CertificateId: "65a1764f33000021008e143e",
		Profile:       "webra-centralized",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	template.KeyType = "rsa-2048"
	request, err := client.NewEnrollRequest(WebRAEnrollRequestParams{
		Profile:  "webra-centralized",
		Template: template,
		Password: "challengepassword",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(request.Password.Value)
}

// Sign a CSR
func TestDecentralizedEnroll(t *testing.T) {
	var csrPem = `-----BEGIN CERTIFICATE REQUEST-----
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
-----END CERTIFICATE REQUEST-----`

	template, err := client.GetEnrollTemplate(WebRAEnrollTemplateParams{
		Csr:     csrPem,
		Profile: "webra-decentralized",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	template.KeyType = "rsa-3072"
	request, err := client.NewEnrollRequest(WebRAEnrollRequestParams{
		Profile:  "webra-decentralized",
		Template: template,
		Password: "challengepassword",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(request.Password.Value)
}

func TestGetEnrollRequest(t *testing.T) {
	request, err := client.GetEnrollRequest("65a27f9b3300003f008e1582")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(request)
}

func TestRevokeRequest(t *testing.T) {
	var csrPem = `-----BEGIN CERTIFICATE-----
MIIEWTCCAkGgAwIBAgIQVskkZfpj966LygN5QqpKfDANBgkqhkiG9w0BAQsFADBD
MQswCQYDVQQGEwJGUjESMBAGA1UEChMJRXZlclRydXN0MSAwHgYDVQQDExdFdmVy
VHJ1c3QgUUEgSXNzdWluZyBDQTAeFw0yNDAxMTMxMjEzMzVaFw0yODAxMTMxMjEz
MzVaMBcxFTATBgNVBAMMDGV2ZXJ0cnVzdC5mcjCCASIwDQYJKoZIhvcNAQEBBQAD
ggEPADCCAQoCggEBAOedT216+7lFZBMLWXT8Ue6vBL9AHMOcirHVVsPAy36wU9oj
V7Tae8o/qITMNTfhRvbH9I2H+NayWam97fI0MKMmRH1M1L1hSqJgB0BHGHjnaeOp
I7UJv1xlgFfFJsSxwRZtUIVDONebjW8KGcuqM2A0p/9f6Q5Otc6Oeh1J6F8Bbhix
nU9iARYZfRPUEF9xxAXfAnoe6cuTz8R+3MMi9CB7yCLU3Uq2kuaDtIx15WfD4bxB
F2XxbcnYqdLFcZfYinih4gqnRD8i2RUn90+S1AO2C4t5Y/+UBdc5jQjwGKtkVVgG
ymQ3qjPYNUbqxKmVc4paYEvRnk2b6Doibn4CPRkCAwEAAaN1MHMwDAYDVR0TAQH/
BAIwADAdBgNVHQ4EFgQUvfD0qxtnSpDKI1vORu7hHbOQCw4wHwYDVR0jBBgwFoAU
FBDcsDMJ96AamspqCNYWV3IlwKkwDgYDVR0PAQH/BAQDAgeAMBMGA1UdJQQMMAoG
CCsGAQUFBwMCMA0GCSqGSIb3DQEBCwUAA4ICAQAYvEvHaMN/Q3Cd06cgmjtTnKXF
lIi+ptNSJw1CboPSBmljMaiyOC/36gcbo4vsD/dj761P5BZ1cmshxbX1MprmN37E
WGFARej6JezAaxg3ULBRmXQ1SP2ov8Ai7Qm4nP/K0hqAUWMGPKKKmBHfCYYle4Gv
m+qPla4qgcSpnhltTPtG4VNdn9V4F9KeTdCgMwWlSVPYMO8G7u21LDfV4nMB/Llc
n0MptbFKxCYnfdWsXPa3Xbr/aT6l6JQi3Tpd2GNYT7I2iTWBE6IOFezyyrEPk0x4
vrjrk5c2UqfspwSmvu2sc9516vz072fKK0cq5Xh8f2sIWlKvK+fM5SQCuSI9KS1m
2psTGYEZ6w/b83unxTXYpPMhnDeBT8sH7JhkX4QOP2RfdWAS75Lvp53D9TKW0Zod
zwqqM9UGlobKqzCPZjF650SABns9i7YAImf5GcPRWoljxmJsmnqmZDC3FVBkkj7x
FjbIHoNzfLzS9fWl7q8ISvl9ogqKwUiV+fZDLuN93TegEqoBAVTJIiycFnaDyH54
ppA6kWKWGyvQUm3f3JO9PURCzJITdZtAm7WNaeVvLjKsLQai3AsENwiRh1tlNbEZ
XrSDPEPo/Sr442rFNJPTJb6pzcipPl5uZKzIxC097vay+jJ9XT308oEf6rqzkDDW
jxB4kE0fnjsnyV/Lzg==
-----END CERTIFICATE-----`
	request, err := client.NewRevokeRequest(WebRARevokeRequestParams{
		CertificatePEM:   csrPem,
		RevocationReason: certificates.Unspecified,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(request)
}

func TestUpdateRequest(t *testing.T) {
	template, err := client.GetUpdateTemplate(WebRAUpdateTemplateParams{
		CertificateId: "65a27f593300001a008e1548",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	if template.Team.Editable {
		template.Team.Value = &HorizonString{"frontend"}
	}
	if template.ContactEmail.Editable {
		template.ContactEmail.Value = &HorizonString{"abcd@free.fr"}
	}

	request, err := client.NewUpdateRequest(WebRAUpdateRequestParams{
		CertificateId: "65a27f593300001a008e1548",
		Template:      template,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(request)
}

func TestMigrateRequest(t *testing.T) {
	targetProfile := "webra-centralized"
	template, err := client.GetMigrateTemplate(WebRAMigrateTemplateParams{
		CertificateId: "65a27f593300001a008e1548",
		Profile:       targetProfile,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(template.Team)
	if template.Team.Editable {
		template.Team.Value = &HorizonString{"backend"}
	}
	if template.ContactEmail.Editable {
		template.ContactEmail.Value = Delete
	}

	request, err := client.NewMigrateRequest(WebRAMigrateRequestParams{
		CertificateId: "65a27f593300001a008e1548",
		Template:      template,
		Profile:       targetProfile,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(request)
}

func TestTemplateAndScepChallenge(t *testing.T) {
	profile := "SCEP_Client"
	template, err := client.GetScepChallengeTemplate(ScepChallengeTemplateParams{
		Profile: profile,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(template)

	template.Subject[0].Value = "moncn"
	template.Sans[0].Value = []string{"monsan.com"}

	requestParams := ScepChallengeRequestParams{
		Profile:  profile,
		Template: template,
	}
	if template.IsDnWhitelist() {
		t.Log("DN Whitelist enabled")
		requestParams.Dn = "CN=abcd,O=efgh"
	}
	request, err := client.NewScepChallengeRequest(requestParams)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(request.Challenge.Value)
}

func TestTemplateAndEstChallenge(t *testing.T) {
	profile := "est-challenge"
	template, err := client.GetEstChallengeTemplate(EstChallengeTemplateParams{
		Profile: profile,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(template)

	template.Subject[0].Value = "moncn"
	//template.Sans[0].Value = []string{"monsan.com"}

	requestParams := EstChallengeRequestParams{
		Profile:  profile,
		Template: template,
	}
	if template.IsDnWhitelist() {
		t.Log("DN Whitelist enabled")
		requestParams.Dn = "CN=abcd,O=efgh"
	}
	request, err := client.NewEstChallengeRequest(requestParams)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(request.Challenge.Value)
}

func TestRecover(t *testing.T) {
	requestParams := WebRARecoverRequestParams{
		CertificateId: "65a17a8f33000020008e14c5",
		Contact:       "toto@toto.com",
		Password:      "monp12",
	}
	request, err := client.NewRecoverRequest(requestParams)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(request.Password.Value)
	t.Log(request.Pkcs12.Value)
}
