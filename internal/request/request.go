package request

import (
	"fmt"
	"io"
	"strings"
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

	parseRequestLine(string(b))
	return &Request{}, nil
}

func parseRequestLine(request string) (*Request, error) {
	lines := strings.Split(request, "\r\n")
	for idx, line := range lines {
		fmt.Printf("%d: %s\n", idx, line)
	}
	parts := strings.Split(lines[0], " ")
	for idx, part := range parts {
		part = strings.TrimSpace(part)
		fmt.Printf("%d: %s\n", idx, part)
	}
	return &Request{}, nil
}
