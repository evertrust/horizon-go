package http

import (
	"encoding/json"
	"mime"
	"strings"

	"gopkg.in/resty.v1"
)

type HorizonResponse struct {
	RestyResponse *resty.Response
}

func (r *HorizonResponse) HasContentType(mimeType string) bool {
	for _, v := range strings.Split(r.RestyResponse.RawResponse.Header.Get("Content-Type"), ",") {
		t, _, err := mime.ParseMediaType(v)
		if err != nil {
			break
		}
		if t == mimeType {
			return true
		}
	}
	return false
}

func (r *HorizonResponse) Json() *json.Decoder {
	return json.NewDecoder(strings.NewReader(string(r.RestyResponse.Body())))
}
