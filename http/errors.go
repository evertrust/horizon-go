package http

import "fmt"

type HorizonErrorResponse struct {
	Code    string `json:"error"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func (e *HorizonErrorResponse) Error() string {
	return fmt.Sprintf("Horizon returned a %s error: %s", e.Code, e.Message)
}
