package fetch

import (
	"io/ioutil"
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
	response, err := Do(args)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(response.Body)
}

// Do fetch do
func Do(args *Request) (*http.Response, error) {
	client := NewClient(args).Do()
	request, err := NewFetch(args).Request()
	if err != nil {
		return nil, err
	}
	return NewFetch(args).Response(request, client)
}
