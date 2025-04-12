package main

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/service_discovery_and_load_balancing/internal/pkg/http"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/service_discovery_and_load_balancing/internal/pkg/http/mechanism"
)

func main() {
	sdConfig := mechanism.ServiceDiscoverMechanismConfig{
		HostIPsMap: map[string][]string{"localhost.com": {"127.0.0.1:8080", "127.0.0.1:8081", "127.0.0.1:8082"}},
	}
	blConfig := mechanism.BlackListMechanismConfig{
		Hosts: []string{"blacklist.com"},
	}

	// 「服務探索 → 負載平衡 → 黑名單」
	client := mechanism.WithServiceDiscoverMechanism(&sdConfig,
		mechanism.WithLoadBalanceMechanism(
			mechanism.WithBlackListMechanism(&blConfig,
				http.NewFakeHttpClient()),
		),
	)
	request := http.NewRequest("http://localhost.com/api/v1/mail")
	client.SendRequest(http.GET, request)

	fmt.Println()
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println()

	// 「黑名單 → 負載平衡 → 服務探索」
	client = mechanism.WithBlackListMechanism(&blConfig,
		mechanism.WithLoadBalanceMechanism(
			mechanism.WithServiceDiscoverMechanism(&sdConfig,
				http.NewFakeHttpClient()),
		),
	)
	// request = http.NewRequest("http://blacklist.com/api/v1/mail")
	client.SendRequest(http.GET, request)
}
