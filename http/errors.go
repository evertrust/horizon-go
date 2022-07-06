package http

import "fmt"

type HorizonErrorResponse struct {
	Code    string `json:"error"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func (e *HorizonErrorResponse) Error() string {
	msg := fmt.Sprintf("Horizon returned a %s error: %s", e.Code, e.Message)
	if e.Detail != "" {
		msg = fmt.Sprintf("%s (%s)", msg, e.Detail)
	}
	return msg
}
