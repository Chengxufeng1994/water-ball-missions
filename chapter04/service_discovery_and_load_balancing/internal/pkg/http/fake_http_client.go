package http

import "fmt"

// Concrete component
type FakeHttpClient struct{}

var _ HttpClient = (*FakeHttpClient)(nil)

func NewFakeHttpClient() *FakeHttpClient {
	return &FakeHttpClient{}
}

func (f *FakeHttpClient) SendRequest(method Method, request Request) (Response, error) {
	fmt.Println("[fake http client] method:", method, "request:", request)
	return Response{}, nil
}
