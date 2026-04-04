package response

import (
	"fmt"

	"github.com/tnaums/httpfromtcp/internal/headers"
)

type StatusCode int

const (
	StatusCodeSuccess             StatusCode = 200
	StatusCodeBadRequest          StatusCode = 400
	StatusCodeInternalServerError StatusCode = 500
)


func getStatusLine(statusCode StatusCode) []byte {
	reasonPhrase := ""
	switch statusCode {
	case StatusCodeSuccess:
		reasonPhrase = "OK"
	case StatusCodeBadRequest:
		reasonPhrase = "Bad Request"
	case StatusCodeInternalServerError:
		reasonPhrase = "Internal Server Error"
	}
	return []byte(fmt.Sprintf("HTTP/1.1 %d %s\r\n", statusCode, reasonPhrase))
}

func GetDefaultHeaders(contentLen int) headers.Headers {
	h := headers.NewHeaders()
	h.Set("Content-Length", fmt.Sprintf("%d", contentLen))
	h.Set("Connection", "close")
	h.Set("Content-Type", "text/plain")
	return h
}

