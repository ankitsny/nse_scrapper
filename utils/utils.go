package utils

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
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

// ExecParallelResponse :
type ExecParallelResponse struct {
	Data  interface{}
	Error error
	Index int
}

// ExecParallel :
func ExecParallel(cLimit int, funcs ...func() interface{}) []ExecParallelResponse {
	semaphoreChan := make(chan struct{}, cLimit)
	resultsChan := make(chan *ExecParallelResponse)

	defer func() {
		close(semaphoreChan)
		close(resultsChan)
	}()

	for i, f := range funcs {
		go func(i int, f func() interface{}) {

			semaphoreChan <- struct{}{}

			data := f()

			r := &ExecParallelResponse{}
			r.Data = data
			r.Index = i

			// write into the channel
			resultsChan <- r

			// free the buffered channel
			<-semaphoreChan

		}(i, f)
	}

	results := make([]ExecParallelResponse, 0)

	for {
		r := <-resultsChan
		results = append(results, *r)

		if len(results) == len(funcs) {
			break
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Index < results[j].Index
	})

	return results
}

// ParallelGETResult :
type ParallelGETResult struct {
	Index    int
	Response *http.Response
	Err      error
}

// ParallelGET :
func ParallelGET(urls []string, query url.Values, cLimit int) []ParallelGETResult {
	semaphoreChan := make(chan struct{}, cLimit)

	resultsChan := make(chan *ParallelGETResult)

	defer func() {
		close(semaphoreChan)
		close(resultsChan)
	}()

	for i, url := range urls {
		go func(i int, url string) {

			// write to semaphore
			semaphoreChan <- struct{}{}

			req := HTTPRequest{}
			req.URL = url
			req.Query = query

			resp, err := req.Do(http.MethodGet)

			pResp := &ParallelGETResult{}

			pResp.Err = err
			pResp.Response = resp
			pResp.Index = i

			resultsChan <- pResp
			// read from chan, just to allow waiting go routine to execute
			<-semaphoreChan

		}(i, url)
	}

	var results []ParallelGETResult
	for {
		resp := <-resultsChan
		results = append(results, *resp)

		// break the loop after completion of all request
		if (len(urls)) == len(results) {
			break
		}
	}

	// Sort the response
	sort.Slice(results, func(i, j int) bool {
		return results[i].Index < results[j].Index
	})

	return results
}

// FilterStr :
func FilterStr(s []string, f func(string) bool) []string {
	var res = make([]string, 0)
	for _, v := range s {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

// MapStr :
func MapStr(s []string, f func(str string) string) []string {
	var res = make([]string, 0)
	for _, v := range s {
		res = append(res, f(v))
	}
	return res
}

// CopyHeader :
func CopyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
