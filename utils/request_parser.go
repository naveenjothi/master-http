package utils

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/naveenjothi/master-http/dto"
)

func ParseRequest(raw []byte) (*dto.Request, error) {
	fmt.Printf("Parsing Request...\n")
	fmt.Printf("Raw bytes: %q\n", raw)

	pos := bytes.Index(raw, []byte("\r\n"))

	if pos == -1 {
		return nil, errors.New("Out of range")
	}

	firstLine := string(raw[:pos])

	values := strings.Split(firstLine, " ")
	if len(values) < 3 {
		return nil, errors.New("400 Bad Request: malformed request line")
	}

	method := values[0]
	uri := values[1]
	proto := values[2]

	pathWithQueryParams := strings.Split(uri, "?")
	path := pathWithQueryParams[0]

	queryParams := make(map[string][]string)

	if len(pathWithQueryParams) > 1 {
		paramsAsArray := strings.Split(pathWithQueryParams[1], "&")

		for _, val := range paramsAsArray {
			keyValues := strings.Split(val, "=")

			if len(keyValues) == 2 {
				key := keyValues[0]
				value := keyValues[1]
				queryParams[key] = append(queryParams[key], value)
			}
		}
	}

	remainingByte := raw[pos+2:]

	headers, err := extractHeaders(remainingByte)
	if err != nil {
		return nil, err
	}

	return &dto.Request{
		Proto:       proto,
		Method:      method,
		Path:        path,
		QueryParams: queryParams,
		Headers:     headers,
	}, nil
}

func extractHeaders(data []byte) (map[string][]string, error) {
	// Headers end at the blank line \r\n\r\n
	end := bytes.Index(data, []byte("\r\n\r\n"))
	if end == -1 {
		// No body; headers go to the end
		end = len(data)
	}

	headers := make(map[string][]string)
	lines := bytes.Split(data[:end], []byte("\r\n"))

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		idx := bytes.IndexByte(line, ':')
		if idx == -1 {
			return nil, fmt.Errorf("400 Bad Request: malformed header: %q", line)
		}
		key := strings.TrimSpace(string(line[:idx]))
		value := strings.TrimSpace(string(line[idx+1:]))
		headers[key] = append(headers[key], value)
	}

	return headers, nil
}
