package headers

import (
	"bytes"
	"fmt"
)

type Headers map[string]string

func NewHeaders() Headers {
	return Headers{}
}

const crlf = "\r\n"

func (h Headers) Parse(data []byte) (n int, done bool, err error) {
	// first check if line starts with crlf; if so, headers are done
	if bytes.HasPrefix(data, []byte(crlf)) {
		return 0, true, nil
	}
	fmt.Println(data)
	
	// if in headers, see if complete line is present
	idx := bytes.Index(data, []byte(crlf))
	if idx == -1 {
		// did not find crlf, nothing to parse
		return 0, false, nil
	}

	// parse the first line
	line := data[:idx]
	fmt.Println(line)
	n = idx + 2 // number of bytes parsed

	// split into key and value
	idx = bytes.Index(line, []byte(":"))
	key := data[:idx]
	if len(key) != len(bytes.TrimSpace(key)) {
		return 0, false, fmt.Errorf("poorly formatted key: %s", string(key))
	}
	value := string(bytes.TrimSpace(data[idx + 1:]))
	h[string(key)] = value

	return n, false, nil
}
