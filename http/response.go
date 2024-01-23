package http

import (
	"bytes"
	"encoding/json"
	"io"
	"mime"
	gohttp "net/http"
	"strings"
)

type HorizonResponse struct {
	HttpResponse *gohttp.Response
	body         []byte
}

func (r *HorizonResponse) HasContentType(mimeType string) bool {
	for _, v := range strings.Split(r.HttpResponse.Header.Get("Content-Type"), ",") {
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
	if r.body == nil {
		// Saving the body since it is a stream
		r.body, _ = io.ReadAll(r.HttpResponse.Body)
	}
	return json.NewDecoder(bytes.NewReader(r.body))
}
