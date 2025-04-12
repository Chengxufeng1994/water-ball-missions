package mechanism

import "github.com/Chengxufeng1994/water-ball-missions/chapter04/service_discovery_and_load_balancing/internal/pkg/http"

type HttpClientMechanismDecorator struct {
	wrapped http.HttpClient
}

var _ http.HttpClient = (*HttpClientMechanismDecorator)(nil)

func newHttpClientMechanism(wrapped http.HttpClient) *HttpClientMechanismDecorator {
	return &HttpClientMechanismDecorator{wrapped: wrapped}
}

func (decorator *HttpClientMechanismDecorator) SendRequest(method http.Method, request http.Request) (http.Response, error) {
	return decorator.wrapped.SendRequest(method, request)
}
