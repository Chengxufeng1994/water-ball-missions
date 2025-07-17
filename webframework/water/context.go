package water

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Chengxufeng1994/water-ball-missions/webframework/water/serdes"
)

type H map[string]interface{}

type Context struct {
	// registry
	registry serdes.Registry
	// request info
	HttpRequest *HttpRequest
	// response info
	HttpResponse *HttpResponse
	// middleware
	handlers []HttpHandlerFunc
	index    int
	// err
	Err error
}

func newContext(
	registry serdes.Registry,
	request *HttpRequest,
	response *HttpResponse,
) *Context {
	return &Context{
		registry:     registry,
		HttpRequest:  request,
		HttpResponse: response,
		handlers:     make([]HttpHandlerFunc, 0),
		index:        -1,
	}
}

func (c *Context) Bind(key string) (any, error) {
	contentType := c.HttpRequest.Req.Header.Get("Content-Type")
	body, err := io.ReadAll(c.HttpRequest.Req.Body)
	if err != nil {
		return nil, err
	}
	return c.registry.Deserialize(contentType, key, body)
}

func (c *Context) PostForm(key string) string {
	return c.HttpRequest.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.HttpRequest.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.HttpResponse.StatusCode = code
	c.HttpResponse.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.HttpResponse.Writer.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...any) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.HttpResponse.Writer.Write(fmt.Appendf(nil, format, values...))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.HttpResponse.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.HttpResponse.Writer, err.Error(), int(StatusCodeInternalServerError))
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.HttpResponse.Writer.Write(data)
}

func (c *Context) Render(code int, contentType string, v any) {
	data, err := c.registry.Serialize(contentType, v)
	if err != nil {
		http.Error(c.HttpResponse.Writer, err.Error(), int(StatusCodeInternalServerError))
		return
	}
	c.SetHeader("Content-Type", contentType)
	c.Status(code)
	if _, err := c.HttpResponse.Writer.Write(data); err != nil {
		// TODO: log error
	}
}

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}
