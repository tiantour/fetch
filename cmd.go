package fetch

import (
	"net/http"
)

// Request request
type Request struct {
	Method string
	URL    string
	Body   []byte
	Header http.Header
	Cert   string
	Key    string
	IP     string
}

// Cmd fetch command
func Cmd(args *Request) ([]byte, error) {
	client := NewClient(args).Do()
	return NewFetch(args).Do(client)
}
