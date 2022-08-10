package http

import (
	"encoding/json"
	"mime"
	"net/http"
	"strings"
)

type HorizonResponse struct {
	BaseResponse *http.Response
}

func (r *HorizonResponse) HasContentType(mimeType string) bool {
	for _, v := range strings.Split(r.BaseResponse.Header.Get("Content-Type"), ",") {
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
	return json.NewDecoder(r.BaseResponse.Body)
}
