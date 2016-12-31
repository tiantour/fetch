package fetch

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	null    struct{}
	methods = map[string]struct{}{
		"GET":    null,
		"POST":   null,
		"PUT":    null,
		"DELETE": null,
		"PATCH":  null,
	}
)

// Cmd fetch command
// date 2016-12-31
// author andy.jiang
func Cmd(method, url string, args ...interface{}) ([]byte, error) {
	method = strings.ToUpper(method)
	if _, ok := methods[method]; !ok {
		return nil, errors.New("method not allowed")
	}
	return operate(method, url, args...)
}

// operate set request
// date 2016-12-31
// author andy.jiang
func operate(method, url string, args ...interface{}) ([]byte, error) {
	// get args
	var body []byte
	var header http.Header
	for _, v := range args {
		switch v.(type) {
		case []byte:
			// set body
			body = v.([]byte)
		case http.Header:
			// set header
			header = v.(http.Header)
		}
	}
	client := &http.Client{}
	// set request
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, nil
	}
	req.Close = true
	req.Header = header
	// get response
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
