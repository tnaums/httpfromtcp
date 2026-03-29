package request

import (
	"bytes"
	"fmt"
	"io"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	requestLine, err := parseRequestLine(b)
	if err != nil {
		return nil, err
	}
	return &Request{
		RequestLine: requestLine,
	}, nil
}

func parseRequestLine(b []byte) (RequestLine, error) {
	lines := bytes.Split(b, []byte("\r\n"))
	parts := bytes.Split(lines[0], []byte(" "))

	// request line must have three parts
	if len(parts) != 3 {
		return RequestLine{}, fmt.Errorf("Requestline did not contain three parts")
	}
	
	// Method can only be uppercase letters
	if checkASCII(parts[0]) == false {
		return RequestLine{}, fmt.Errorf("Method contains invalid character")
	}

	// only HTTP 1.1 is supported
	version := bytes.TrimPrefix(parts[2], []byte("HTTP/"))
	if string(version) != "1.1" {
		return RequestLine{}, fmt.Errorf("Only HTTP/1.1 is supported.")
	}
	return RequestLine{
		HttpVersion: string(version),
		RequestTarget: string(parts[1]),
		Method: string(parts[0]),
	}, nil
}

func checkASCII(b []byte) bool {
	for _, r := range b {
		if r > 90 || r < 65 {
			return false
		}
	}
	return true
}
