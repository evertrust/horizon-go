package certificateprofiles

import "github.com/evertrust/horizon-go/http"

type Client struct {
	Http *http.Client
}

func (c *Client) Get(id string) (*Profile, error) {
	response, err := c.Http.Get("/api/v1/certificate/profiles/" + id)
	if err != nil {
		return nil, err
	}
	var profile Profile
	err = response.Json().Decode(&profile)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}
