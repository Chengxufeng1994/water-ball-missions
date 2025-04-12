package mechanism

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/service_discovery_and_load_balancing/internal/pkg/http"
)

type ServiceDiscoverMechanismConfig struct {
	HostIPsMap map[string][]string
	FailCache  map[string]map[string]time.Time // host => ip => failTime
	Mutex      sync.RWMutex
	TTL        time.Duration // 例如：10 分鐘
}

// Concrete Decorator
type ServiceDiscoverMechanism struct {
	*ServiceDiscoverMechanismConfig
	*HttpClientMechanismDecorator
}

var _ http.HttpClient = (*ServiceDiscoverMechanism)(nil)

func WithServiceDiscoverMechanism(config *ServiceDiscoverMechanismConfig, wrapped http.HttpClient) http.HttpClient {
	if config.FailCache == nil {
		config.FailCache = make(map[string]map[string]time.Time)
	}
	if config.TTL == 0 {
		config.TTL = 10 * time.Minute
	}

	sdm := &ServiceDiscoverMechanism{
		ServiceDiscoverMechanismConfig: config,
		HttpClientMechanismDecorator:   newHttpClientMechanism(wrapped),
	}

	go sdm.tick()
	return sdm
}

func (decorator *ServiceDiscoverMechanism) tick() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now()
		decorator.Mutex.Lock()
		for host, failedIPs := range decorator.FailCache {
			for ip, failTime := range failedIPs {
				if now.Sub(failTime) > decorator.TTL {
					delete(decorator.FailCache[host], ip)
				}
			}
		}
		decorator.Mutex.Unlock()
	}
}

func (decorator *ServiceDiscoverMechanism) SendRequest(method http.Method, request http.Request) (http.Response, error) {
	fmt.Println("[service discover mechanism]")
	decorator.Mutex.RLock()
	ips, ok := decorator.HostIPsMap[request.Host]
	failCache := decorator.FailCache[request.Host]
	decorator.Mutex.RUnlock()

	if !ok {
		fmt.Println("[service discover mechanism] host not found")
		return http.Response{}, errors.New("host not found")
	}

	now := time.Now()
	filtered := make([]string, 0, len(ips))

	for _, ip := range ips {
		failTime, failed := failCache[ip]
		if !failed || now.Sub(failTime) > decorator.TTL {
			filtered = append(filtered, ip)
		}
	}

	if len(filtered) != 0 {
		fmt.Println("[service discover mechanism] all IPs unhealthy")
		request.IPs = filtered
	}

	resp, err := decorator.wrapped.SendRequest(method, request)
	if err != nil {
		fmt.Println("[service discover mechanism] error:", err)
		decorator.Mutex.Lock()
		failCache[request.TargetUrl] = now
		decorator.Mutex.Unlock()
		return resp, err
	}

	return resp, nil
}
