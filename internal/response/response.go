package response

import (
	"fmt"
	"io"

	"github.com/tnaums/httpfromtcp/internal/headers"
)

type StatusCode int

//type Response struct {
//}

const (
	StatusOK                  StatusCode = 200
	StatusBadRequest          StatusCode = 400
	StatusInternalServerError StatusCode = 500
)

func WriteStatusLine(w io.Writer, statusCode StatusCode) error {
	out := []byte{}
	switch statusCode {
	case StatusOK:
		out = []byte("HTTP/1.1 200 OK\r\n")
	case StatusBadRequest:
		out = []byte("HTTP/1.1 400 Bad Request\r\n")
	case StatusInternalServerError:
		out = []byte("HTTP/1.1 500 Internal Server Error\r\n")
	default:
		out = []byte("HTTP/1.1 \r\n ")
	}
	w.Write(out)
	return nil
}

func GetDefaultHeaders(contentLen int) headers.Headers {
	returnHeaders := headers.NewHeaders()
	returnHeaders["Content-Length"] = fmt.Sprintf("%d", contentLen)
	returnHeaders["Connection"] = "close"
	returnHeaders["Content-Type"] = "text/plain"
	return returnHeaders
}

func WriteHeaders(w io.Writer, headers headers.Headers) error {
	for k, v := range headers {
		w.Write([]byte(fmt.Sprintf("%s: %s\r\n", k, v)))
	}
	w.Write([]byte("\r\n"))
	return nil
}
