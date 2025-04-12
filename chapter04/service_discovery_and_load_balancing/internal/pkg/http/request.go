package http

import (
	stdurl "net/url"
)

type Request struct {
	Scheme      string
	Host        string
	Path        string
	Query       map[string][]string
	IPs         []string
	OriginalUrl string
	TargetUrl   string
}

func NewRequest(url string) Request {
	u, _ := stdurl.Parse(url)

	query := make(map[string][]string)
	for k, v := range u.Query() {
		query[k] = v
	}

	return Request{
		Scheme:      u.Scheme,
		Host:        u.Host,
		Path:        u.Path,
		Query:       query,
		IPs:         []string{u.Host},
		OriginalUrl: url,
		TargetUrl:   url,
	}
}
