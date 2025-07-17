package water

import "net/http"

type HttpRequest struct {
	Req    *http.Request
	Path   string
	Method Method
	Params map[string]string
}
