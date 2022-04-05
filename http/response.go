package http

import (
	"encoding/json"
	"net/http"
)

type HorizonResponse struct {
	BaseResponse *http.Response
}

func (r *HorizonResponse) Json() *json.Decoder {
	return json.NewDecoder(r.BaseResponse.Body)
}
