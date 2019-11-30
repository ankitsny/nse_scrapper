package utils

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// ToBytes :
func ToBytes(data interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := gob.NewEncoder(buffer)
	if err := encoder.Encode(data); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// HTTPRequest : Request payload
type HTTPRequest struct {
	Body  interface{}
	Query url.Values
	URL   string
}

var client *http.Client

func getHTTPClient() *http.Client {
	if client == nil {
		client = &http.Client{
			Timeout: time.Second * 25,
		}
	}
	return client
}

func makeHTTPRequest(method string, requestParams *HTTPRequest) (*http.Response, error) {

	url, err := url.Parse(requestParams.URL)

	if err != nil {
		return nil, err
	}
	// build http request
	req := &http.Request{
		Method: method,
		URL:    url,
	}

	// build query params
	if requestParams.Query != nil {
		req.URL.RawQuery = requestParams.Query.Encode()
	}

	req.Header = http.Header{}

	// set body, if exist
	if method != http.MethodGet && requestParams.Body != nil {
		bbyt, err := ToBytes(requestParams.Body)
		if err != nil {
			return nil, err
		}
		req.Body = ioutil.NopCloser(bytes.NewReader(bbyt))
		defer req.Body.Close()
	}
	fmt.Println("CALLING", req.URL.String())
	return getHTTPClient().Do(req)
}

// Do : call method
// accepts http method
func (r HTTPRequest) Do(method string) (*http.Response, error) {
	return makeHTTPRequest(method, &r)
}
