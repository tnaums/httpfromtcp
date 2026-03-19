package request

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode"
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
		return &Request{}, err
	}

	rp, err := parseRequestLine(string(b))
	if err != nil {
		return &Request{}, err
	}
	fmt.Printf("HttpVersion: %s\n\n", rp.RequestLine.HttpVersion)
	fmt.Printf("RequestTarget: %s\n\n", rp.RequestLine.RequestTarget)
	fmt.Printf("Method: %s\n\n", rp.RequestLine.Method)
	return rp, nil
}

func parseRequestLine(request string) (*Request, error) {
	lines := strings.Split(request, "\r\n")
	for idx, line := range lines {
		fmt.Printf("%d: %s\n", idx, line)
	}
	parts := strings.Split(lines[0], " ")
	for _, part := range parts {
		part = strings.TrimSpace(part)
	}
	if len(parts) != 3 {
		err := errors.New("Request Line has incorrect number of parts.")
		return &Request{}, err
	}
	Method := parts[0]
	for _, r := range Method {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			err := errors.New("Method must be all uppercase letters.")
			return &Request{}, err
		}
	}
	RequestTarget := parts[1]
	HttpVersion := parts[2]
	HttpVersion = strings.TrimPrefix(HttpVersion, "HTTP/")
	assembled := RequestLine{
		HttpVersion:   HttpVersion,
		RequestTarget: RequestTarget,
		Method:        Method,
	}
	return &Request{
		RequestLine: assembled,
	}, nil
}
