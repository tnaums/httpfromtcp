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
		return 2, true, nil
	}
	
	// if in headers, see if complete line is present
	idx := bytes.Index(data, []byte(crlf))
	if idx == -1 {
		// did not find crlf, nothing to parse
		return 0, false, nil
	}

	// parse the first line
	line := data[:idx]
	n = idx + 2 // number of bytes parsed

	// split into key and value
	idx = bytes.Index(line, []byte(":"))
	key := data[:idx]
	if len(key) != len(bytes.TrimRight(key, " ")) {
		return 0, false, fmt.Errorf("poorly formatted key: %s", string(key))
	}
	key = bytes.TrimLeft(key, " ")
	key = bytes.ToLower(key)
	keyString := string(key)
	for _, r := range keyString {
		if isInvalidChar(r) {
			return 0, false, fmt.Errorf("invalid character in field name: %s", string(r))
		}
	}
	value := string(bytes.TrimSpace(data[idx + 1:]))
	h[keyString] = value

	return n, false, nil
}


func isInvalidChar(r rune) bool {
    // return true if r is NOT in the allowed set
    isLetter := (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
	isDigit := r >= '0' && r <= '9'
	isSetA := r == '!'
	isSetB := r >= '#' && r <= '\''
	isSetC := r == '*' || r == '+'
	isSetD := r == '-' || r == '.'
	isSetE := r >= '^' && r <= '`'
	isSetF := r == '|' || r == '~'
    return !isLetter && !isDigit && !isSetA && !isSetB && !isSetC && !isSetD && !isSetE && !isSetF
}
