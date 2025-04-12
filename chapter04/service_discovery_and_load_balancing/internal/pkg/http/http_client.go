package http

type HttpClient interface {
	SendRequest(method Method, request Request) (Response, error)
}
