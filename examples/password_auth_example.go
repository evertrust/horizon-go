package examples

import (
	horizon "github.com/evertrust/horizon-go/v2"
)

// ExamplePasswordAuth demonstrates how to set up a client with password authentication.
func ExamplePasswordAuth() *horizon.APIClient {
	c := horizon.NewConfiguration()
	c.Servers[0].URL = "https://horizon.example.com"
	c.SetPasswordAuth("user", "password")
	client := horizon.NewAPIClient(c)
	return client
}
