package fetch

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
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
func Cmd(args Request) ([]byte, error) {
	client := &http.Client{}
	// set tls
	if args.Cert != "" && args.Key != "" {
		client = setTLS(args)
	}
	// set proxy
	if args.IP != "" {
		client = setProxy(args)
	}
	// set request
	req, err := http.NewRequest(args.Method, args.URL, bytes.NewReader(args.Body))
	if err != nil {
		return nil, nil
	}
	req.Close = true
	req.Header = args.Header
	// get response
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// setTLS set http tls
func setTLS(args Request) *http.Client {
	cert, err := tls.LoadX509KeyPair(args.Cert, args.Key)
	if err != nil {
		return nil
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{
			cert,
		},
	}
	transport := &http.Transport{
		TLSClientConfig: config,
	}
	return &http.Client{
		Transport: transport,
	}
}

// setProxy set http proxy
func setProxy(args Request) *http.Client {
	proxy := func(*http.Request) (*url.URL, error) {
		return url.Parse(args.IP)
	}
	transport := &http.Transport{
		Proxy: proxy,
	}
	return &http.Client{
		Transport: transport,
	}
}
