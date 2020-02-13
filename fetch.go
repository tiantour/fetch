package fetch

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// Fetch Fetch
type Fetch struct {
	Method string
	URL    string
	Body   []byte
	Header http.Header
}

// NewFetch new fetch
func NewFetch(args *Request) *Fetch {
	return &Fetch{
		Method: args.Method,
		URL:    args.URL,
		Body:   args.Body,
		Header: args.Header,
	}
}

// Do fetch do
func (f *Fetch) Do(args *http.Client) ([]byte, error) {
	request, err := f.Request()
	if err != nil {
		return nil, err
	}
	return f.Response(request, args)
}

// Request request
func (f *Fetch) Request() (*http.Request, error) {
	request, err := http.NewRequest(f.Method, f.URL, bytes.NewReader(f.Body))
	if err != nil {
		return nil, err
	}
	request.Close = true
	request.Header = f.Header
	return request, nil
}

// Response response
func (f *Fetch) Response(args *http.Request, client *http.Client) ([]byte, error) {
	response, err := client.Do(args)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}
