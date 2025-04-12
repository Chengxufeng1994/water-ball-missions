package mechanism

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/service_discovery_and_load_balancing/internal/pkg/http"
)

// Concrete Decorator
type LoadBalanceMechanismDecorator struct {
	*HttpClientMechanismDecorator
	selectedIndex int
}

var _ http.HttpClient = (*LoadBalanceMechanismDecorator)(nil)

func WithLoadBalanceMechanism(wrapped http.HttpClient) http.HttpClient {
	return &LoadBalanceMechanismDecorator{
		HttpClientMechanismDecorator: newHttpClientMechanism(wrapped),
		selectedIndex:                0,
	}
}

func (decorator *LoadBalanceMechanismDecorator) SendRequest(requestMethod http.Method, request http.Request) (http.Response, error) {
	fmt.Println("[load balance mechanism]")
	if len(request.IPs) == 0 {
		return decorator.wrapped.SendRequest(requestMethod, request)
	}

	targetIP := request.IPs[decorator.selectedIndex]
	request.TargetUrl = fmt.Sprintf("%s://%s%s", request.Scheme, targetIP, request.Path)

	decorator.selectedIndex = (decorator.selectedIndex + 1) % len(request.IPs)

	return decorator.wrapped.SendRequest(requestMethod, request)
}
