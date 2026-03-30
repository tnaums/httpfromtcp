package request

import (
	"bytes"
	"fmt"
	"io"
)

type RequestState int

const (
	initialized RequestState = iota
	done
)

type Request struct {
	RequestLine RequestLine
	RequestState RequestState
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
	requestLine, n, err := parseRequestLine(b)
	if err != nil {
		return nil, err
	}

	if n == 0 {
		return nil, nil
	}
	return &Request{
		RequestLine: *requestLine,
	}, nil
}

func parseRequestLine(b []byte) (*RequestLine, int, error) {
	lines := bytes.Split(b, []byte("\r\n"))
	parts := bytes.Split(lines[0], []byte(" "))
	n := len(lines[0]) + 2
	
	// request line must have three parts
	if len(parts) != 3 {
		return nil, 0, fmt.Errorf("Requestline did not contain three parts")
	}
	
	// Method can only be uppercase letters
	if checkASCII(parts[0]) == false {
		return nil, 0, fmt.Errorf("Method contains invalid character")
	}

	// only HTTP 1.1 is supported
	version := bytes.TrimPrefix(parts[2], []byte("HTTP/"))
	if len(version) == len(parts[2]) {
		return nil, 0, fmt.Errorf("malformed HTTP version: %s", parts[2])
	}
	if string(version) != "1.1" {
		return nil, 0, fmt.Errorf("Only HTTP/1.1 is supported: %s", parts[2])
	}
	return &RequestLine{
		HttpVersion: string(version),
		RequestTarget: string(parts[1]),
		Method: string(parts[0]),
	}, n, nil
}

func checkASCII(b []byte) bool {
	for _, r := range b {
		if r > 'Z' || r < 'A' {
			return false
		}
	}
	return true
}
