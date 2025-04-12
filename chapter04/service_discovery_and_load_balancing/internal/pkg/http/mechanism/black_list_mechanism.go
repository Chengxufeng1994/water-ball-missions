package mechanism

import (
	"errors"
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/service_discovery_and_load_balancing/internal/pkg/http"
)

type BlackListMechanismConfig struct {
	Hosts []string
}

type BlackListMechanismDecorator struct {
	*BlackListMechanismConfig
	*HttpClientMechanismDecorator
}

func WithBlackListMechanism(config *BlackListMechanismConfig, wrapped http.HttpClient) http.HttpClient {
	return &BlackListMechanismDecorator{
		BlackListMechanismConfig:     config,
		HttpClientMechanismDecorator: newHttpClientMechanism(wrapped),
	}
}

func (b *BlackListMechanismDecorator) SendRequest(method http.Method, request http.Request) (http.Response, error) {
	fmt.Println("[black list mechanism] ")
	for _, host := range b.Hosts {
		if host == request.Host {
			return http.Response{}, errors.New("in black list")
		}
	}

	return b.wrapped.SendRequest(method, request)
}
