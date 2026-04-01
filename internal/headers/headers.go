package headers

import (
	"bytes"
	"fmt"
	"strings"
)

type Headers map[string]string

func NewHeaders() Headers {
	return map[string]string{}
}

func (h Headers) Parse(data []byte) (int, bool, error) {
	n := bytes.Index(data, []byte("\r\n"))

	// not enough data; \r\n not found
	if n == -1 {
		return 0, false, nil
	}
	// \r\n found at start of data; done parsing headers
	if n == 0 {
		return 2, true, nil
	}
	
	parts := bytes.SplitN(data[:n], []byte(":"), 2)
	key := string(parts[0])
	if key != strings.TrimRight(key, " ") {
		return 0, false, fmt.Errorf("invalid header name: %s", key)
	}

	value := bytes.TrimSpace(parts[1])
	key = strings.TrimSpace(key)
	
	h.Set(key, string(value))
	return n + 2, false, nil
}

func (h Headers) Set(key, value string) {
	h[key] = value
}
