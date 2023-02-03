package http

import "fmt"

type HorizonErrorResponse struct {
	Code    string `json:"error"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

type HorizonMultipleErrorsResponse []HorizonErrorResponse

func (e *HorizonErrorResponse) Error() string {
	msg := fmt.Sprintf("Horizon returned a %s error: %s", e.Code, e.Message)
	if e.Detail != "" {
		msg = fmt.Sprintf("%s (%s)", msg, e.Detail)
	}
	return msg
}

func (e *HorizonMultipleErrorsResponse) Error() string {
	msg := "Horizon returned multiple errors:\n"
	for _, err := range *e {
		toAppend := fmt.Sprintf("\t- %s: %s\n", err.Code, err.Message)
		if err.Detail != "" {
			toAppend = fmt.Sprintf("\t- %s: %s (%s)\n", err.Code, err.Message, err.Detail)
		}
		msg = fmt.Sprintf("%s%s", msg, toAppend)
	}
	return msg
}

type Feature int

const (
	TrustchainDecoding Feature = iota
)

func (feature Feature) String() string {
	switch feature {
	case TrustchainDecoding:
		return "trustchains decoding"
	}
	return "unknown feature"
}

type NotImplementedError struct {
	Feature       Feature `json:"feature"`
	ImplementedIn string  `json:"implementedIn"`
}

func (e *NotImplementedError) Error() string {
	return fmt.Sprintf("The current Horizon version doesn't support this feature (%s). Please upgrade the instance to a version >= %s", e.Feature.String(), e.ImplementedIn)
}
