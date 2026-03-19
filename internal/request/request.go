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

const crlf = "\r\n"

// RequestFromReader reads an HTTP request from an io.Reader and returns a Request struct.
func RequestFromReader(reader io.Reader) (*Request, error) {
	req, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	// Parse the request line
	requestLine, err := parseRequestLine(req)
	if err != nil {
		return nil, err
	}
	return &Request{
		RequestLine: *requestLine,
	}, nil
}

// parseRequestLine parses the request line of an HTTP request and returns a RequestLine struct.
func parseRequestLine(data []byte) (*RequestLine, error) {
	idx := bytes.Index(data, []byte(crlf))
	if idx == -1 {
		return nil, fmt.Errorf("invalid request: no CRLF found")
	}
	requestLineText := string(data[:idx])

}
