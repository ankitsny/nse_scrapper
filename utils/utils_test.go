package utils

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestToBytes(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToBytes(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHTTPClient(t *testing.T) {
	tests := []struct {
		name string
		want *http.Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHTTPClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getHTTPClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeHTTPRequest(t *testing.T) {
	type args struct {
		method        string
		requestParams *HTTPRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := makeHTTPRequest(tt.args.method, tt.args.requestParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("makeHTTPRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeHTTPRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPRequest_Do(t *testing.T) {
	type args struct {
		method string
	}
	tests := []struct {
		name    string
		r       HTTPRequest
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Do(tt.args.method)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPRequest.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPRequest.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExecParallel(t *testing.T) {
	type args struct {
		cLimit int
		funcs  []func() interface{}
	}
	tests := []struct {
		name string
		args args
		want []ExecParallelResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExecParallel(tt.args.cLimit, tt.args.funcs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExecParallel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParallelGET(t *testing.T) {
	type args struct {
		urls   []string
		query  url.Values
		cLimit int
	}
	tests := []struct {
		name string
		args args
		want []ParallelGETResult
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParallelGET(tt.args.urls, tt.args.query, tt.args.cLimit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParallelGET() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterStr(t *testing.T) {
	type args struct {
		s []string
		f func(string) bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterStr(tt.args.s, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapStr(t *testing.T) {
	type args struct {
		s []string
		f func(str string) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapStr(tt.args.s, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopyHeader(t *testing.T) {
	type args struct {
		dst http.Header
		src http.Header
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CopyHeader(tt.args.dst, tt.args.src)
		})
	}
}
