package dto

type Request struct {
	Method      string              `json:"method"`
	Path        string              `json:"path"`
	Headers     map[string][]string `json:"headers"`
	QueryParams map[string][]string `json:"queryParams"`
	Body        []byte              `json:"body"`
	Proto       string              `json:"proto"`
}
