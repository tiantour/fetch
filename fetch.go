package fetch

import (
	"bytes"
	"crypto/tls"
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
}

// Cmd fetch command
func Cmd(args Request) ([]byte, error) {
	client := &http.Client{}
	// set tls
	if args.Cert != "" && args.Key != "" {
		cert, err := tls.LoadX509KeyPair(args.Cert, args.Key)
		if err != nil {
			return nil, err
		}

		config := &tls.Config{
			Certificates: []tls.Certificate{
				cert,
			},
		}

		tr := &http.Transport{
			TLSClientConfig: config,
		}

		client = &http.Client{
			Transport: tr,
		}
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
