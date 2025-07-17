package water

import "net/http"

type HttpResponse struct {
	Writer     http.ResponseWriter
	StatusCode int
}
