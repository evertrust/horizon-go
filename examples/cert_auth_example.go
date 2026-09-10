package examples

import (
	"crypto/tls"
	"log"

	horizon "github.com/evertrust/horizon-go/v2"
)

// ExampleCertAuth demonstrates how to set up a client with certificate authentication.
func ExampleCertAuth() *horizon.APIClient {
	c := horizon.NewConfiguration()
	c.Servers[0].URL = "https://horizon.example.com"

	// Load the certificate
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatal(err)
	}
	c.SetCertAuth(cert)

	client := horizon.NewAPIClient(c)
	return client
}
