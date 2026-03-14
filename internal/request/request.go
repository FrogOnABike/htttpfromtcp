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

// parseRequestLine parses the request line of an HTTP request and returns a RequestLine struct.
func parseRequestLine(line string) (*RequestLine, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid request line: %s", line)
	}
	return &RequestLine{
		Method:        parts[0],
		RequestTarget: parts[1],
		HttpVersion:   parts[2],
	}, nil
}

// RequestFromReader reads an HTTP request from an io.Reader and returns a Request struct.
func RequestFromReader(reader io.Reader) (*Request, error) {
	req, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(req), "\r\n")
	if len(lines) < 1 {
		return nil, fmt.Errorf("invalid request: no lines")
	}

	// Parse the request line
	requestLine, err := parseRequestLine(lines[0])
	if err != nil {
		return nil, err
	}

	// Validate the HTTP method only contains alphabetic characters
	for _, r := range requestLine.Method {
		if !('A' <= r && r <= 'Z') && !('a' <= r && r <= 'z') {
			return nil, fmt.Errorf("invalid HTTP method: %s", requestLine.Method)
		}
	}

	// Validate the HTTP version is 1.1 as that's all we're supporting in this project
	if requestLine.HttpVersion != "HTTP/1.1" {
		return nil, fmt.Errorf("unsupported HTTP version: %s", requestLine.HttpVersion)
	}

	return &Request{
		RequestLine: *requestLine,
	}, nil
}
