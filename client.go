package fetch

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
)

// Client client
type Client struct {
	Cert string
	Key  string
	IP   string
}

// NewClient new client
func NewClient(args *Request) *Client {
	return &Client{
		Cert: args.Cert,
		Key:  args.Key,
		IP:   args.IP,
	}
}

// Do fetch do
func (c *Client) Do() *http.Client {
	client := &http.Client{}
	if c.Cert != "" {
		client.Transport = c.TLS()
	}
	if c.IP != "" {
		client.Transport = c.Proxy()
	}
	return client
}

// TLS fetch tls
func (c *Client) TLS() *http.Transport {
	cert, err := tls.LoadX509KeyPair(c.Cert, c.Key)
	if err != nil {
		log.Fatalf("load cert error: %v", err)
	}
	return &http.Transport{
		TLSClientConfig: &tls.Config{Certificates: []tls.Certificate{
			cert,
		}},
	}
}

// Proxy fetch proxy
func (c *Client) Proxy() *http.Transport {
	return &http.Transport{
		Proxy: func(*http.Request) (*url.URL, error) {
			return url.Parse(c.IP)
		},
	}
}
